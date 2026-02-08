package main

import (
	"fmt"
	"sync"
	"time"
)

type RateLimiter struct {
	counter map[string]int
	mu      sync.Mutex
}

func (rl *RateLimiter) Allow(ip string) bool {
	rl.mu.Lock()
	defer rl.mu.Unlock()

	rl.counter[ip]++
	return rl.counter[ip] <= 5
}

func main() {
	rl := &RateLimiter{}
	rl.counter = make(map[string]int)

	// Simulate a burst of 10 requests from the same IP
	for i := 0; i < 10; i++ {
		go func(id int) {
			if rl.Allow("192.168.1.1") {
				fmt.Printf("Request %d: Allowed\n", id)
			} else {
				fmt.Printf("Request %d: BLOCKED\n", id)
			}
		}(i)
	}

	// Block main to let goroutines finish
	time.Sleep(1 * time.Second)
}
