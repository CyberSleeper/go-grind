package main

import (
	"strconv"
	"sync"
	"testing"
)

func TestLongRunningSimulation(t *testing.T) {
	orc := NewOrchestrator(10000, 10000, 10000, 10000, 5)
	var input = make(chan DataEntry, 10000) // Buffer larger so generators don't deadlock

	var wgGenerators sync.WaitGroup
	var wgProducers sync.WaitGroup
	var wgItems sync.WaitGroup

	for i := 0; i < orc.ProducerCnt; i++ {
		wgProducers.Add(1)
		go func(id int) {
			defer wgProducers.Done()
			for {
				data, ok := <-input
				if !ok {
					return
				}
				data.producerId = id
				orc.Producers[id].Enqueue(data)
			}
		}(i)
	}

	// Send 10,000 items dynamically
	for i := range 10000 {
		wgGenerators.Add(1)
		wgItems.Add(1)
		go func(val int) {
			defer wgGenerators.Done()
			curData := DataEntry{
				value:      strconv.Itoa(val),
				producerId: -1, // Just to mark unassigned value
				consumerId: -1, // Just to mark unassigned value
			}
			input <- curData
		}(i)
	}

	for i := 0; i < orc.ConsumerCnt; i++ {
		go func(id int) {
			orc.Consumers[id].Start(&wgItems)
		}(i)
	}

	// 1. Wait for generator loops to finish
	wgGenerators.Wait()

	// 2. Shut off the first conveyer belt
	close(input)

	// 3. Wait for Producers to completely drain input and shut down
	wgProducers.Wait()

	// 4. Shut off the main conveyer belt
	// close(orc.Stream.Stream)

	// 5. Wait for Consumers to completely drain the main stream and shut down
	wgItems.Wait()

	// Print out the statistics
	t.Log("== Load Balancing Statistics ==")
	for i, v := range orc.Producers {
		if v.processedDataCnt > 1 {
			t.Logf("Producer %d processed: %d", i, v.processedDataCnt)
		}
	}
	for i, v := range orc.Consumers {
		if v.processedDataCnt > 1 {
			t.Logf("Consumer %d processed: %d", i, v.processedDataCnt)
		}
	}
}

// BenchmarkQueue runs a highly accurate statistical test on the
// performance of the queue. We use the Finite Batch pattern here
func BenchmarkQueue(b *testing.B) {
	for n := 0; n < b.N; n++ {
		var wgProducers sync.WaitGroup
		var wgConsumers sync.WaitGroup
		orc := NewOrchestrator(200, 10000, 100, 100, 5)

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
