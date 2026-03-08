package main

import (
	"sync"
	"testing"
)

// 1. Basic FIFO Test
func TestRingBuffer_Basic(t *testing.T) {
	rb := NewRingBuffer(4)

	// Enqueue 2 items
	if !rb.Enqueue(Order{ID: 1}) {
		t.Errorf("Expected Enqueue to succeed")
	}
	if !rb.Enqueue(Order{ID: 2}) {
		t.Errorf("Expected Enqueue to succeed")
	}

	// Dequeue 1st item
	order, ok := rb.Dequeue()
	if !ok || order.ID != 1 {
		t.Errorf("Expected to dequeue Order 1, got %v (ok: %v)", order.ID, ok)
	}

	// Dequeue 2nd item
	order, ok = rb.Dequeue()
	if !ok || order.ID != 2 {
		t.Errorf("Expected to dequeue Order 2, got %v (ok: %v)", order.ID, ok)
	}
}

// 2. Full and Empty Boundary Test
func TestRingBuffer_Boundaries(t *testing.T) {
	rb := NewRingBuffer(2) // Very small buffer!

	// Fill it up
	rb.Enqueue(Order{ID: 1})
	rb.Enqueue(Order{ID: 2})

	// Try to overfill it (Should Fail safely)
	if rb.Enqueue(Order{ID: 3}) {
		t.Errorf("Expected Enqueue to fail because buffer is full")
	}

	// Empty it out
	rb.Dequeue()
	rb.Dequeue()

	// Try to over-empty it (Should Fail safely)
	_, ok := rb.Dequeue()
	if ok {
		t.Errorf("Expected Dequeue to fail because buffer is empty")
	}
}

// 3. The Wraparound Test (Tests your Bitwise Mask!)
func TestRingBuffer_Wraparound(t *testing.T) {
	rb := NewRingBuffer(4)

	// Push 4, Pop 4 (Cursors are now at index 4, physically out of bounds)
	for i := 1; i <= 4; i++ {
		rb.Enqueue(Order{ID: int64(i)})
	}
	for i := 1; i <= 4; i++ {
		rb.Dequeue()
	}

	// The next push MUST wrap around to index 0 using your bitwise mask
	if !rb.Enqueue(Order{ID: 5}) {
		t.Errorf("Failed to wrap around and enqueue!")
	}

	order, ok := rb.Dequeue()
	if !ok || order.ID != 5 {
		t.Errorf("Failed to dequeue wrapped order! Got ID: %d", order.ID)
	}
}

// 4. Concurrency Test (Single Producer, Single Consumer)
func TestRingBuffer_ConcurrentSPSC(t *testing.T) {
	rb := NewRingBuffer(1024)
	var wg sync.WaitGroup

	totalMessages := 100_000

	wg.Add(2)

	// Producer Goroutine (The Sequencer)
	go func() {
		defer wg.Done()
		for i := 1; i <= totalMessages; i++ {
			// Spin-wait if the buffer is full (simulating high load)
			for !rb.Enqueue(Order{ID: int64(i)}) {
			}
		}
	}()

	// Consumer Goroutine (The Matching Engine)
	go func() {
		defer wg.Done()
		count := 0
		for count < totalMessages {
			order, ok := rb.Dequeue()
			if ok {
				count++
				// Ensure strict First-In-First-Out ordering
				if order.ID != int64(count) {
					t.Errorf("Race condition detected! Expected ID %d, got %d", count, order.ID)
				}
			}
		}
	}()

	wg.Wait()
}
