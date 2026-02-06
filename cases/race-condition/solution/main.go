package main

import (
	"fmt"
	"sync"
)

var count int

func main() {
	var wg sync.WaitGroup
	mu := sync.Mutex{}
	for range 1000 {
		wg.Go(func() {
			mu.Lock()
			count++
			mu.Unlock()
		})
	}
	wg.Wait()
	fmt.Println(count)
}
