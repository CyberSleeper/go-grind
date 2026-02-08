package main

import (
	"fmt"
)

// Stage 1: Convert numbers to a channel stream
func gen(nums ...int) <-chan int {
	out := make(chan int)
	go func() {
		for _, v := range nums {
			out <- v
		}
		close(out)
	}()
	return out
}

// Stage 2: Square the numbers
func sq(in <-chan int) <-chan int {
	out := make(chan int)
	go func() {
		for v := range in {
			out <- v * v
		}
		close(out)
	}()
	return out
}

func main() {
	// Set up the pipeline
	// 1. Generate numbers 2 and 3
	c := gen(2, 3, 4, 5)

	// 2. Square them
	out := sq(c)

	// 3. Consume the output
	fmt.Println("Pipeline Results:")
	for n := range out {
		fmt.Println(n)
	}
}
