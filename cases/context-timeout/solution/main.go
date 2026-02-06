package main

import (
	"context"
	"fmt"
	"time"
)

func slowOperation(ctx context.Context) {
	select {
	case <-time.After(5 * time.Second):
		fmt.Println("Work Done")
	case <-ctx.Done():
		fmt.Println("Timeout!")
	}
}

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()

	fmt.Println("Starting operation...")
	slowOperation(ctx)
	fmt.Println("Main finished")
}
