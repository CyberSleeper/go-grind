package main

import (
	"sync"
	"sync/atomic"
	"testing"
)

// TestOrderSequencer_Sequential verifies that IDs are strictly increasing
// when called sequentially on a single thread.
func TestOrderSequencer_Sequential(t *testing.T) {
	seq := NewOrderSequencer()
	var prevID int64 = -1

	// Generate 100,000 IDs quickly to trigger the sequence increment logic
	for i := 0; i < 100000; i++ {
		id := seq.NextID()

		if id <= prevID {
			t.Fatalf("IDs are not strictly monotonic! Previous: %d, Current: %d", prevID, id)
		}
		prevID = id
	}
}

// TestOrderSequencer_Concurrent unleashes absolute chaos to ensure
// no two goroutines ever receive the same ID.
func TestOrderSequencer_Concurrent(t *testing.T) {
	seq := NewOrderSequencer()

	const numGoroutines = 500
	const opsPerGoroutine = 1000

	var wg sync.WaitGroup
	wg.Add(numGoroutines)

	seen := sync.Map{}

	var duplicateCount atomic.Int64

	for range numGoroutines {
		go func() {
			defer wg.Done()

			for j := 0; j < opsPerGoroutine; j++ {
				id := seq.NextID()

				if _, loaded := seen.LoadOrStore(id, true); loaded {
					duplicateCount.Add(1)
				}
			}
		}()
	}

	wg.Wait()

	if duplicateCount.Load() > 0 {
		t.Fatalf("CRITICAL FAILURE: Found %d duplicate IDs under concurrent load!", duplicateCount.Load())
	}
}
