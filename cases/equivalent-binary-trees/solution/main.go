package main

import (
	"fmt"

	"golang.org/x/tour/tree"
)

// Walk walks the tree t sending all values
// from the tree to the channel ch.
func Walk(t *tree.Tree, ch chan int) {
	recWalk(t, ch)
	close(ch)
}

func recWalk(t *tree.Tree, ch chan int) {
	if t == nil {
		return
	}

	if t.Left != nil {
		recWalk(t.Left, ch)
	}
	ch <- t.Value
	if t.Right != nil {
		recWalk(t.Right, ch)
	}
}

// Same determines whether the trees
// t1 and t2 contain the same values.
func Same(t1, t2 *tree.Tree) bool {
	ch1 := make(chan int)
	ch2 := make(chan int)

	go Walk(t1, ch1)
	go Walk(t2, ch2)

	for {
		v1, ok1 := <-ch1
		v2, ok2 := <-ch2

		if ok1 != ok2 {
			return false
		}
		if !ok1 {
			return true
		}
		if v1 != v2 {
			return false
		}
	}
}

func main() {
	ch := make(chan int)
	go Walk(tree.New(1), ch)

	for v := range ch {
		fmt.Println(v)
	}

	fmt.Println(Same(tree.New(1), tree.New(1)))
	fmt.Println(Same(tree.New(1), tree.New(2)))
}
