package main

import (
	"fmt"
	"sync"
	"time"
)

func fetchUser() {
	time.Sleep(100 * time.Millisecond)
	fmt.Println("User fetched")
}

func fetchPortfolio() {
	time.Sleep(300 * time.Millisecond)
	fmt.Println("Portfolio fetched")
}

func main() {
	var wg sync.WaitGroup

	wg.Add(1)
	go func() {
		defer wg.Done()
		fetchUser()
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		fetchPortfolio()
	}()

	wg.Wait()

	fmt.Println("All Done")
}
