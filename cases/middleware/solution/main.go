package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

func LoggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		next.ServeHTTP(w, r)
		log.Printf("[%s] %s took %dms", r.Method, r.URL.Path, time.Since(start).Milliseconds())
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
