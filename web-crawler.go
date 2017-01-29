package main

import (
	"fmt"
)

type Fetcher interface {
	// Fetch returns the body of URL and
	// a slice of URLs found on that page.
	Fetch(url string) (body string, urls []string, err error)
}

var fetchedUrls = make(map[string]bool)

type FetchedUrl struct {
	url string
	body string
}

// Crawl uses fetcher to recursively crawl
// pages starting with url, to a maximum of depth.
func Crawl(url string, depth int, fetcher Fetcher, ret chan FetchedUrl) {
	defer close(ret)
	if depth <= 0 {
		return
	}

	body, urls, err := fetcher.Fetch(url)
	if err != nil {
        	ret <- FetchedUrl{}
		return
	}

	ret <- FetchedUrl{body:body,url:url}

	result := make([]chan FetchedUrl, 0)
	for _, u := range urls {
		if fetchedUrls[u] == false {
			fetchedUrls[u] = true
			c1 := make(chan FetchedUrl)
			result = append(result, c1)
			go Crawl(u, depth-1, fetcher, c1)
		}
	}

	for i := range result {
		for s := range result[i] {
			ret <- s
		}
	}

	return
}

func main() {
	channel := make(chan FetchedUrl)
	go Crawl("http://golang.org/", 4, fetcher, channel)

	for s := range channel {
		fmt.Printf("Found: %s %q\n", s.url, s.body)
	}
}


// fakeFetcher is Fetcher that returns canned results.
type fakeFetcher map[string]*fakeResult

type fakeResult struct {
	body string
	urls []string
}

func (f fakeFetcher) Fetch(url string) (string, []string, error) {
	if res, ok := f[url]; ok {
		return res.body, res.urls, nil
	}
	return "", nil, fmt.Errorf("not found: %s", url)
}

// fetcher is a populated fakeFetcher.
var fetcher = fakeFetcher{
	"http://golang.org/": &fakeResult{
		"The Go Programming Language",
		[]string{
			"http://golang.org/pkg/",
			"http://golang.org/cmd/",
		},
	},
	"http://golang.org/pkg/": &fakeResult{
		"Packages",
		[]string{
			"http://golang.org/",
			"http://golang.org/cmd/",
			"http://golang.org/pkg/fmt/",
			"http://golang.org/pkg/os/",
		},
	},
	"http://golang.org/pkg/fmt/": &fakeResult{
		"Package fmt",
		[]string{
			"http://golang.org/",
			"http://golang.org/pkg/",
		},
	},
	"http://golang.org/pkg/os/": &fakeResult{
		"Package os",
		[]string{
			"http://golang.org/",
			"http://golang.org/pkg/",
		},
	},
}

