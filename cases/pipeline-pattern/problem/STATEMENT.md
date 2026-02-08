## Problem 9: The Pipeline Pattern (Channels)

### 📌 Concept

A **Pipeline** is a series of stages connected by channels, where each stage is a group of goroutines running the same function.

**Benefits:**

- **Separation of Concerns:** Each function does one thing well.
- **Stream Processing:** You don't need to load all 1GB of data into memory; you process it line-by-line as it flows through.

### 📝 Task

Implement a 2-stage pipeline.

1.  **`gen(nums ...int) <-chan int`**:
    - Takes a list of integers.
    - Starts a goroutine to send them to a channel.
    - Closes the channel when done.
    - Returns the channel (Receive-Only).
2.  **`sq(in <-chan int) <-chan int`**:
    - Takes an input channel.
    - Starts a goroutine to read from `in`, square the number (`n*n`), and send it to an output channel.
    - Closes the output channel when `in` is closed.
    - Returns the output channel.
3.  **`main`**:
    - Set up the pipeline: `c := gen(2, 3)` -> `out := sq(c)`.
    - Range over `out` and print the results.

### 🚫 Starter Code

```go
package main

import (
    "fmt"
)

// Stage 1: Convert numbers to a channel stream
func gen(nums ...int) <-chan int {
    out := make(chan int)
    go func() {
        // TODO: Send nums to 'out', then close 'out'
    }()
    return out
}

// Stage 2: Square the numbers
func sq(in <-chan int) <-chan int {
    out := make(chan int)
    go func() {
        // TODO: Read from 'in', square it, send to 'out', then close 'out'
    }()
    return out
}

func main() {
    // Set up the pipeline
    // 1. Generate numbers 2 and 3
    c := gen(2, 3)

    // 2. Square them
    out := sq(c)

    // 3. Consume the output
    fmt.Println("Pipeline Results:")
    for n := range out {
        fmt.Println(n)
    }
}
```
