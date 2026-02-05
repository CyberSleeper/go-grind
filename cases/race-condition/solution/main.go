package main

import (
	"fmt"
	"sync"
)

var count int

func main() {
	var wg sync.WaitGroup
	for range 1000 {
		wg.Go(func() {
			// CRITICAL SECTION START
			count++ // <--- DATA RACE!
			// CRITICAL SECTION END
		})
	}
	wg.Wait()
	fmt.Println(count) // Output is unpredictable (e.g., 950 instead of 1000)
}
