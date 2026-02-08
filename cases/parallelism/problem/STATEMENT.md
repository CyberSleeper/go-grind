## Problem 10: Fan-Out, Fan-In (Parallelism)

### 📌 Concept

In the previous pipeline, the `sq` stage was sequential. If squaring a number took 1 second, processing 10 numbers would take 10 seconds.

**The Solution:**

1.  **Fan-Out:** Start multiple workers (instances of `sq`) to read from the _same_ input channel. They share the workload automatically.
2.  **Fan-In:** Create a `merge` function that takes multiple result channels and combines them back into a single output channel for the main function to consume.

### 📝 Task

Implement the `merge` function to combine results from multiple channels.

1.  **`merge(cs ...<-chan int) <-chan int`**:
    - **Input:** A slice of input channels.
    - **Output:** A single channel containing all values from all inputs.
    - **Logic:**
      - Create a `sync.WaitGroup`.
      - For each input channel in `cs`, start a goroutine that copies values from that input channel to the single output channel.
      - **Crucial:** Start _another_ separate goroutine that waits for all workers to finish (`wg.Wait()`) and then closes the output channel. This ensures `main` doesn't hang.

### 🚫 Starter Code

```go
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
        // TODO
    }

    // 2. Start a goroutine for each input channel
    wg.Add(len(cs))
    for _, c := range cs {
        go output(c)
    }

    // 3. Start a monitor goroutine to close 'out' once everyone is done
    go func() {
        // TODO: Wait and Close
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
```
