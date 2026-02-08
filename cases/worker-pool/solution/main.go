package main

import (
	"fmt"
	"sync"
	"time"
)

const WORKER_COUNT = 3
const JOB_COUNT = 5

func worker(id int, jobs <-chan int) {
	for j := range jobs {
		time.Sleep(1 * time.Second)
		fmt.Printf("Worker %d finished job %d\n", id, j)
	}
}

func main() {
	jobs := make(chan int, 100)
	var wg sync.WaitGroup

	for i := range WORKER_COUNT {
		wg.Go(func() {
			worker(i, jobs)
		})
	}

	for i := range JOB_COUNT {
		jobs <- i
	}
	close(jobs)

	wg.Wait()
}
