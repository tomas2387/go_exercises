package main

import (
	"golang.org/x/tour/tree"
	"fmt"
)

func doWalk (t *tree.Tree, ch chan int) {
	if (t == nil) {
		return
	}
	doWalk(t.Left, ch)
	ch <- t.Value
	doWalk(t.Right, ch)
}


// Walk walks the tree t sending all values
// from the tree to the channel ch.
func Walk(t *tree.Tree, ch chan int) {
	doWalk(t, ch)
	close(ch)
}

// Same determines whether the trees
// t1 and t2 contain the same values.
func Same(t1, t2 *tree.Tree) bool {
	c1, c2 := make(chan int, 10), make(chan int, 10)
	go Walk(t1, c1)
	go Walk(t2, c2)

	for i := 0; i < 10; i++ {
		x1, ok1 := <- c1
		x2, ok2 := <- c2
		if !ok1 && !ok2 {
			return true
		} else if !ok1 && ok2 || ok1 && !ok2 {
			fmt.Println("Channel closed before the other!")
			return false
		} else if  x1 != x2 {
			fmt.Printf("%v != %v \n", x1, x2)
			return false
		}
	}

	return true
}

func main() {
	fmt.Println(Same(tree.New(1), tree.New(1)))
	fmt.Println(Same(tree.New(1), tree.New(2)))
}

