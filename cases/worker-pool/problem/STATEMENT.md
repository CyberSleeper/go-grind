## Problem 6: The Worker Pool (Buffered Channels)

### 📌 Concept

Spawning a goroutine for every single task is dangerous. If you have 1 million tasks, you will crash the runtime or overwhelm your database.

A **Worker Pool** limits concurrency. You start a fixed number of workers (e.g., 3) that pull tasks from a shared channel.

### 📝 Task

1.  **The Job:** A simple integer (`id`).
2.  **The Worker:** A function that takes `id`, sleeps for 1 second, and prints "Worker X finished job Y".
3.  **The Main:**
    - Create a `jobs` channel.
    - Start **3 worker goroutines**.
    - Send **5 jobs** (integers 1-5) to the channel.
    - **Close** the `jobs` channel.
    - Wait for all workers to finish.

### 🚫 Starter Code

```go
package main

import (
    "fmt"
    "sync"
    "time"
)

func worker(id int, jobs <-chan int, wg *sync.WaitGroup) {
    // TODO:
    // 1. Loop over 'jobs' channel until it closes.
    // 2. Simulate work (Sleep 1s).
    // 3. Print "Worker [id] finished job [job_id]"
    // 4. Call wg.Done() ONLY when the worker exits (not per job).
}

func main() {
    jobs := make(chan int, 100)
    var wg sync.WaitGroup

    // TODO:
    // 1. Start 3 workers
    // 2. Send 5 jobs
    // 3. Close channel so workers know to stop
    // 4. Wait
}
```
