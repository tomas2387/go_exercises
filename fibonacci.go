package main

import "fmt"

func fibRecursive(i int) int {
	if i == 0 || i == 1 {
		return i
	} else if i == 2 {
		return 1
	}

	return fibRecursive(i - 1) + fibRecursive(i - 2)
}

// fibonacci is a function that returns
// a function that returns an int.
func fibonacci() func() int {
	i := 0
	return func() int {
		defer (func() {
			i += 1
		})()
		return fibRecursive(i)
	}
}

func main() {
	f := fibonacci()
	for i := 0; i < 10; i++ {
		fmt.Println(f())
	}
}
