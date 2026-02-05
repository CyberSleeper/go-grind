## Problem 2: The WaitGroup (Basic Concurrency)

### 📌 Concept

Golang is famous for concurrency. However, simply using `go func()` spawns a thread that runs in the background. If the main function finishes before the background thread, the program exits immediately, and the background work is killed.

To coordinate multiple goroutines, we use `sync.WaitGroup`. It acts like a counter:

1.  **Add(N):** "I am starting N tasks."
2.  **Done():** "One task finished." (Decrements counter)
3.  **Wait():** "Block here until counter becomes 0."

### 📝 Task

Refactor the code below to run `fetchUser` and `fetchPortfolio` **concurrently** (at the same time).
The main function **must block** and wait for both to finish before printing "All Done".

### 🚫 Starter Code

```go
package main

import (
    "fmt"
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
    // CURRENT STATUS: Sequential (Slow)
    // TODO: Refactor to run these in parallel using goroutines & WaitGroup

    fetchUser()
    fetchPortfolio()

    fmt.Println("All Done")
}
```
