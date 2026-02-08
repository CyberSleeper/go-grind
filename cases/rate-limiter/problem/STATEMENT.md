## Problem 8: Rate Limiter (Thread Safety)

### 📌 Concept

Rate Limiting protects your application from abuse (DDoS, Brute Force attacks) by rejecting requests after a certain threshold.

Since HTTP requests happen in parallel goroutines, your Rate Limiter **must be thread-safe**. If two requests come from the same IP at the exact same nanosecond, you must count **both** of them correctly.

### 📝 Task

Implement a `RateLimiter` struct.

1.  **State:** It needs a `map[string]int` to track visit counts per IP.
2.  **Safety:** It needs a `sync.Mutex` to protect that map.
3.  **Logic:** Implement `Allow(ip string) bool`.
    - Lock.
    - Increment the count for that IP.
    - If count > 5, return `false` (Blocked).
    - Else, return `true` (Allowed).
    - Unlock.

### 🚫 Starter Code

```go
package main

import (
    "fmt"
    "sync"
    "time"
)

type RateLimiter struct {
    // TODO: Add Mutex and Map
}

func (rl *RateLimiter) Allow(ip string) bool {
    // TODO:
    // 1. Thread-safe access
    // 2. Increment count
    // 3. Return true if count <= 5, else false
    return true
}

func main() {
    rl := &RateLimiter{}

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
```
