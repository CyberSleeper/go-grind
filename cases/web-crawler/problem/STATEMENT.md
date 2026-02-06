## Problem: Concurrency - Web Crawler

### 📌 Concept

This is a classic concurrency problem that tests your ability to manage **shared state** (visited URLs) safely while spawning multiple workers (goroutines) to fetch pages in parallel.

This problem appears in "A Tour of Go" and is highly relevant to backend engineering roles (e.g., building scrapers, indexers, or distributed task processors).

### 📝 Task

Modify the `Crawl` function to fetch URLs **in parallel** without fetching the same URL twice.

**Requirements:**

1.  **Parallelism:** Don't wait for one URL to finish before starting the next. Use `go` routines.
2.  **Synchronization:** Use `sync.WaitGroup` to ensure the main function waits for all goroutines to finish.
3.  **Thread Safety:** Use `sync.Mutex` (or a channel) to protect the `visited` map. If two goroutines try to visit "https://golang.org/" at the same time, only one should succeed.

### 🚫 Starter Code (Sequential & Unsafe)

```go
package main

import (
	"fmt"
)

type Fetcher interface {
	// Fetch returns the body of URL and
	// a slice of URLs found on that page.
	Fetch(url string) (body string, urls []string, err error)
}

// Crawl uses fetcher to recursively crawl
// pages starting with url, to a maximum of depth.
func Crawl(url string, depth int, fetcher Fetcher) {
	// TODO: Fetch URLs in parallel.
	// TODO: Don't fetch the same URL twice.

	if depth <= 0 {
		return
	}

	body, urls, err := fetcher.Fetch(url)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Printf("found: %s %q\n", url, body)

	for _, u := range urls {
		Crawl(u, depth-1, fetcher)
	}
	return
}

func main() {
	Crawl("[https://golang.org/](https://golang.org/)", 4, fetcher)
}
```

### 💡 Hint

You will need a struct (e.g., `SafeCounter` or `Visited`) that holds a `map[string]bool` and a `sync.Mutex`. When checking if a URL is visited, be careful of the **"Check-Then-Act"** race condition. You should lock, check, set, and unlock in one atomic operation.
