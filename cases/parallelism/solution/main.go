package main

import (
	"fmt"
	"sync"
)

// gen converts a list of numbers to a channel
func gen(nums ...int) <-chan int {
	out := make(chan int)
	go func() {
		for _, n := range nums {
			out <- n
		}
		close(out)
	}()
	return out
}

// sq reads from in, squares, and sends to out
func sq(in <-chan int) <-chan int {
	out := make(chan int)
	go func() {
		for n := range in {
			out <- n * n
		}
		close(out)
	}()
	return out
}

func merge(cs ...<-chan int) <-chan int {
	out := make(chan int)
	var wg sync.WaitGroup

	// 1. Define a helper function that copies values from 'c' to 'out'
	//    Don't forget to call wg.Done() when the channel closes!
	output := func(c <-chan int) {
		for v := range c {
			out <- v
		}
		wg.Done()
	}

	// 2. Start a goroutine for each input channel
	wg.Add(len(cs))
	for _, c := range cs {
		go output(c)
	}

	go func() {
		wg.Wait()
		close(out)
	}()

	return out
}

func main() {
	in := gen(2, 3)

	// FAN-OUT: Distribute work to 2 workers
	// Both read from the SAME 'in' channel.
	c1 := sq(in)
	c2 := sq(in)

	// FAN-IN: Collect results from both
	for n := range merge(c1, c2) {
		fmt.Println(n) // Order might be random!
	}
}
