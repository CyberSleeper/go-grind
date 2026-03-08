package main

import "sync/atomic"

// Order represents our standard payload
type Order struct {
	ID       int64
	Quantity int64
	Price    int64
}

// RingBuffer is a lock-free SPSC queue
type RingBuffer struct {
	buffer []Order
	mask   uint64 // Used for fast bitwise modulo

	// These must be manipulated atomically!
	writeCursor uint64
	readCursor  uint64
}

func NewOrder() Order {
	return Order{}
}

// NewRingBuffer creates a new ring buffer.
// Size MUST be a power of 2 (e.g., 1024, 2048, 4096).
func NewRingBuffer(size uint64) *RingBuffer {
	if size&(size-1) != 0 {
		panic("size must be a power of 2")
	}
	return &RingBuffer{
		buffer: make([]Order, size),
		mask:   size - 1,
	}
}

func (rb *RingBuffer) inc(cnt *uint64) {
	atomic.AddUint64(cnt, 1)
}

func (rb *RingBuffer) GetValue(idx uint64) *Order {
	return &rb.buffer[idx&rb.mask]
}

func (rb *RingBuffer) SetValue(idx uint64, val Order) {
	rb.buffer[idx&rb.mask] = val
}

func (rb *RingBuffer) GetCursor() (uint64, uint64) {
	return atomic.LoadUint64(&rb.writeCursor), atomic.LoadUint64(&rb.readCursor)
}

func (rb *RingBuffer) IsFull(wc, rc uint64) bool {
	return wc-rc > rb.mask
}

func (rb *RingBuffer) IsEmpty(wc, rc uint64) bool {
	return wc == rc
}

// Enqueue attempts to push an order into the buffer.
// Returns false if the buffer is completely full.
func (rb *RingBuffer) Enqueue(order Order) bool {
	wc, rc := rb.GetCursor()

	if rb.IsFull(wc, rc) {
		return false
	}
	rb.SetValue(wc, order)
	rb.inc(&rb.writeCursor)

	return true
}

// Dequeue attempts to pop an order from the buffer.
// Returns (Order{}, false) if the buffer is completely empty.
func (rb *RingBuffer) Dequeue() (Order, bool) {
	wc, rc := rb.GetCursor()

	if rb.IsEmpty(wc, rc) {
		return Order{}, false
	}
	curOrder := *rb.GetValue(rc)
	rb.inc(&rb.readCursor)

	return curOrder, true
}
