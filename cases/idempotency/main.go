package main

import (
	"context"
	"sync"
)

type call struct {
	wg    sync.WaitGroup
	value string
	err   error
}

// TradeProcessor ensures that concurrent trade requests with the same
// idempotency key are only executed once.
type TradeProcessor struct {
	Done map[string]*call
	mu   sync.Mutex
}

func NewCall() *call {
	return &call{}
}

func NewTradeProcessor() *TradeProcessor {
	return &TradeProcessor{
		Done: make(map[string]*call),
	}
}

// Execute handles the concurrent trade requests.
// - idempotencyKey: The unique ID for this trade (e.g., "req-999")
// - tradeFunc: The actual API call to the exchange. Must be called ONCE per key.
func (tp *TradeProcessor) Execute(ctx context.Context, idempotencyKey string, tradeFunc func() (string, error)) (string, error) {
	tp.mu.Lock()

	if v, exists := tp.Done[idempotencyKey]; exists {
		tp.mu.Unlock()
		v.wg.Wait()
		return v.value, v.err
	}

	c := NewCall()
	c.wg.Add(1)
	tp.Done[idempotencyKey] = c
	tp.mu.Unlock()

	c.value, c.err = tradeFunc()

	// Delete the idempotency kv to prevent OOM
	// It is safe to do this since all the concurrent request are waiting in v.wg.Wait() line
	tp.mu.Lock()
	delete(tp.Done, idempotencyKey)
	tp.mu.Unlock()

	c.wg.Done()

	return c.value, c.err
}
