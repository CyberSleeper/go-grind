package main

import (
	"sync"
	"time"
)

const curEpoch int64 = 1772809024000000

// OrderSequencer generates strictly increasing, thread-safe IDs
// for the trading matching engine.
type OrderSequencer struct {
	mu       sync.Mutex
	lastTime int64
	seqCnt   int
}

func getNow() int64 {
	return time.Now().UnixMicro() - curEpoch
}

func NewOrderSequencer() *OrderSequencer {
	return &OrderSequencer{
		lastTime: getNow(),
	}
}

func (s *OrderSequencer) waitUntilNextMicro(currentMicro int64) int64 {
	next := getNow()
	// Spin-wait until the clock moves strictly PAST the provided microsecond
	for next <= currentMicro {
		next = getNow()
	}
	return next
}

// NextID must be completely thread-safe and return a strictly
// greater ID than the previous call, even within the same millisecond.
func (s *OrderSequencer) NextID() int64 {
	s.mu.Lock()

	now := getNow()

	if s.seqCnt >= (1<<12)-1 {
		now = s.waitUntilNextMicro(now)
	}

	if s.lastTime > now {
		now = s.waitUntilNextMicro(s.lastTime)
	}

	if s.lastTime == now {
		s.seqCnt++
	} else {
		s.lastTime = now
		s.seqCnt = 0
	}
	nxtId := int64(s.seqCnt)

	s.mu.Unlock()

	// [ 1 unused signed bit ] + [ 51 bits for timestamp ] + [ 12 bits for sequence count ]
	// 51 bits of microsecond will fine until 71 years

	// do modulo to prevent overflow after shl
	now = (now % (int64(1) << 51)) << 12
	nxtId |= now

	return nxtId
}
