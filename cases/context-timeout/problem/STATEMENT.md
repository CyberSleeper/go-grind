## Problem 4: The Context (Timeouts)

### 📌 Concept

In backend systems, a `context.Context` carries deadlines, cancellation signals, and request-scoped values across API boundaries.

**The Scenario:** You are calling a slow external service (like a Database or 3rd party API). If it takes too long (e.g., > 2 seconds), you want to **cancel** the request immediately to free up resources, rather than waiting forever.

### 📝 Task

Modify the `slowOperation` function to accept a `context.Context`.

1.  **Main:** Create a context with a **2-second timeout**. Pass it to `slowOperation`.
2.  **SlowOperation:** Simulate work that takes **5 seconds**.
3.  **The Logic:**
    - If the work finishes before the timeout, print "Work Done".
    - If the context is cancelled (timeout hits) _before_ work finishes, stop immediately and print "Timeout!".
    - **Constraint:** You cannot just use `time.Sleep(5 * time.Second)` because `Sleep` is blocking and cannot be interrupted. You need a way to listen to `ctx.Done()` _while_ waiting.

### 🚫 Starter Code

```go
package main

import (
    "context"
    "fmt"
    "time"
)

func slowOperation(ctx context.Context) {
    // TODO:
    // 1. Simulate 5 seconds of work.
    // 2. BUT, return early if ctx.Done() is closed.

    // Hint: Use 'select' with 'time.After' vs 'ctx.Done()'.
}

func main() {
    // TODO: Create a context that times out after 2 seconds
    ctx := context.TODO()

    fmt.Println("Starting operation...")
    slowOperation(ctx)
    fmt.Println("Main finished")
}
```

### 💡 Hint

You cannot interrupt a standard `time.Sleep`. Instead, think about "waiting for a timer channel" inside a `select` block.
