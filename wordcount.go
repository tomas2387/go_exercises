package main

import (
	"golang.org/x/tour/wc"
	"strings"
)

func WordCount(s string) (result map[string]int)  {
	result = make(map[string]int)
	for _, value := range strings.Fields(s) {
		result[value] += 1
	}
	return
}

func main() {
	wc.Test(WordCount)
}

