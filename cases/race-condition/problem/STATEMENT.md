## Problem 3: The Race Condition (Mutex)

### 📌 Concept

A **Data Race** occurs when multiple goroutines try to read and write to the same memory location at the same time.

In the code below, `count++` looks like one line, but to the CPU, it is three separate operations:

1.  **Read** current value of `count`.
2.  **Increment** the value.
3.  **Write** the new value back to memory.

If two goroutines do step 1 at the exact same time, they will both see the same old value (e.g., 5), increment it to 6, and write 6. We essentially "lost" one count.

### 📝 Task

Fix the code below so that `count` reliably equals 1000 every time.
You need to protect the critical section using `sync.Mutex`.

**Requirements:**

1.  Define a `sync.Mutex`.
2.  **Lock** the mutex before accessing `count`.
3.  **Unlock** the mutex immediately after updating `count`.

### 🚫 Starter Code

```go
package main

import (
    "fmt"
    "sync"
)

var count int

func main() {
    var wg sync.WaitGroup
    for i := 0; i < 1000; i++ {
        wg.Add(1)
        go func() {
            defer wg.Done()
            // CRITICAL SECTION START
            count++ // <--- DATA RACE!
            // CRITICAL SECTION END
        }()
    }
    wg.Wait()
    fmt.Println(count) // Output is unpredictable (e.g., 950 instead of 1000)
}
```
