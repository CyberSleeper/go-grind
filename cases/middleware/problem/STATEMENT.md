## Problem 7: HTTP Middleware (Logging)

### 📌 Concept

Middleware is a function that wraps an `http.Handler` to perform actions **before** or **after** the main request processing.

**Common Uses:**

- Logging request duration.
- Authenticating users (checking JWT).
- Handling CORS or GZIP compression.

### 📝 Task

Write a middleware function `LoggingMiddleware`.

1.  **Input:** An `http.Handler` (the "next" handler to call).
2.  **Output:** An `http.Handler` (the wrapped version).
3.  **Logic:**
    - Record `start := time.Now()`.
    - Call `next.ServeHTTP(w, r)` (Process the actual request).
    - Calculate `duration := time.Since(start)`.
    - Print: `Method: [GET], Path: [/hello], Duration: [12ms]`.

### 🚫 Starter Code

```go
package main

import (
    "fmt"
    "log"
    "net/http"
    "time"
)

// TODO: Implement this function
func LoggingMiddleware(next http.Handler) http.Handler {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        // 1. Record start time
        // 2. Call the next handler
        // 3. Log the duration
    })
}

func mainHandler(w http.ResponseWriter, r *http.Request) {
    // Simulate work
    time.Sleep(100 * time.Millisecond)
    w.Write([]byte("Hello, World!"))
}

func main() {
    // Create the core handler
    coreHandler := http.HandlerFunc(mainHandler)

    // Wrap it with middleware
    wrappedHandler := LoggingMiddleware(coreHandler)

    fmt.Println("Server starting on :8080...")
    log.Fatal(http.ListenAndServe(":8080", wrappedHandler))
}
```
