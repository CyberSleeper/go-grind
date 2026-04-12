package main

import (
	"strconv"
	"sync"
	"testing"
	"time"
)

func TestLongRunningSimulation(t *testing.T) {
	orc := NewOrchestrator(50, 10, 10)
	var input = make(chan DataEntry, 10000) // Buffer larger so generators don't deadlock
	var output = make(chan DataEntry, 100)
	var terminate = make(chan int, 1)

	for i := 0; i < orc.ProducerCnt; i++ {
		go func(id int) {
			for {
				data := <-input
				data.producerId = id
				orc.Producers[id].Enqueue(data)
			}
		}(i)
	}

	for i := 0; i < orc.ConsumerCnt; i++ {
		go func(id int) {
			for {
				val, ok := orc.Consumers[id].Dequeue()
				if ok {
					val.consumerId = id
					output <- val
				} else {
					break
				}
			}
		}(i)
	}

	// Send 10,000 items dynamically
	for i := 0; i < 10000; i++ {
		go func(val int) {
			curData := DataEntry{
				value:      strconv.Itoa(val),
				producerId: -1, // Just to mark unassigned value
				consumerId: -1, // Just to mark unassigned value
			}
			input <- curData
		}(i)
	}

	// Background consumer of output to prevent blockage
	go func() {
	Mainloop:
		for {
			select {
			case <-output:
				// Nyeh
			case <-terminate:
				break Mainloop
			}
		}
	}()

	// Wait briefly for the 10k items to be distributed and processed.
	time.Sleep(200 * time.Millisecond)
	terminate <- 1

	// Print out the statistics
	t.Log("== Load Balancing Statistics ==")
	for i, v := range orc.Producers {
		t.Logf("Producer %d processed: %d", i, v.processedDataCnt)
	}
	for i, v := range orc.Consumers {
		t.Logf("Consumer %d processed: %d", i, v.processedDataCnt)
	}
}

// BenchmarkQueue runs a highly accurate statistical test on the
// performance of the queue. We use the Finite Batch pattern here
func BenchmarkQueue(b *testing.B) {
	for n := 0; n < b.N; n++ {
		var wgProducers sync.WaitGroup
		var wgConsumers sync.WaitGroup
		orc := NewOrchestrator(200, 100, 100)

		for i := 0; i < 100; i++ {
			wgProducers.Add(1)
			go func(val int) {
				defer wgProducers.Done()

				// Simulate some small CPU work before producing if you want realistic results
				// time.Sleep(1 * time.Microsecond)

				orc.Producers[0].Enqueue(DataEntry{
					value: strconv.Itoa(val),
				})
			}(i)
		}

		for i := 0; i < orc.ConsumerCnt; i++ {
			wgConsumers.Add(1)
			go func(id int) {
				defer wgConsumers.Done()
				for {
					_, ok := orc.Consumers[id].Dequeue()
					if !ok {
						break
					}
					// Simulate small work
					// time.Sleep(1 * time.Microsecond)
				}
			}(i)
		}

		wgProducers.Wait()
		close(orc.Stream.Stream)
		wgConsumers.Wait()
	}
}
