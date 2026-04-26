package main

import (
	"errors"
	"math/rand"
	"sync"
	"time"
)

// Note: Golang Chan already has builtin lock.
// It only allows one enqueue XOR one dequeue
// at one time

type Producer struct {
	Stream           *DataStream
	processedDataCnt int
}

type Consumer struct {
	Stream           *DataStream
	DeadLetter       *DataStream
	processedDataCnt int
	MaxRetries       int
}

type DataStream struct {
	Stream chan DataEntry
	cap    int
}

type Orchestrator struct {
	Producers   []*Producer
	Consumers   []*Consumer
	Stream      *DataStream
	DeadLetter  *DataStream
	Capacity    int
	ProducerCnt int
	ConsumerCnt int
	MaxRetries  int
}

type DataEntry struct {
	value      string
	producerId int
	consumerId int
	retryCount int
}

func (p *Producer) Enqueue(x DataEntry) {
	p.processedDataCnt++
	p.Stream.Stream <- x
}

func (c *Consumer) Dequeue() (DataEntry, bool) {
	val, ok := <-c.Stream.Stream
	if ok {
		c.processedDataCnt++
	}

	return val, ok
}

func (c *Consumer) Consume(wg *sync.WaitGroup) bool {
	val, ok := c.Dequeue()
	if !ok {
		return false
	}
	if err := ProcessItem(val.value); err != nil {
		val.retryCount++
		if val.retryCount > c.MaxRetries {
			c.DeadLetter.Stream <- val
			wg.Done()
		} else {
			go func() {
				c.Stream.Stream <- val
			}()
		}
	} else {
		wg.Done()
	}
	return true
}

func (c *Consumer) Start(wg *sync.WaitGroup) {
	for {
		c.Consume(wg)
	}
}

// ProcessItem simulates a real-world task that takes time and can fail.
func ProcessItem(value string) error {
	// Simulate work taking between 10ms - 50ms
	time.Sleep(time.Duration(rand.Intn(40)+10) * time.Millisecond)
	chance := rand.Intn(100) // 0 to 99
	if chance < 60 {
		return nil // 60% chance: Immediate Success!
	} else if chance < 90 {
		return errors.New("temporary network failure") // 30% chance: Fails, but might work if retried
	} else {
		return errors.New("fatal data corruption") // 10% chance: Hard fail. Could theoretically send straight to DLQ!
	}
}

func (d *DataStream) Len() int {
	return len(d.Stream)
}

func (d *DataStream) Close() {
	close(d.Stream)
}

func NewDataStream(cap int) *DataStream {
	return &DataStream{
		Stream: make(chan DataEntry, cap),
		cap:    cap,
	}
}

func NewProducer(d *DataStream) *Producer {
	return &Producer{
		Stream: d,
	}
}

func NewConsumer(d, deadLetter *DataStream, maxRetries int) *Consumer {
	return &Consumer{
		Stream:     d,
		DeadLetter: deadLetter,
		MaxRetries: maxRetries,
	}
}

func NewOrchestrator(cap, deadLetterCap, producerCnt, consumerCnt, maxRetries int) *Orchestrator {
	stream := NewDataStream(cap)
	deadLetter := NewDataStream(deadLetterCap)
	producers := make([]*Producer, producerCnt)
	consumers := make([]*Consumer, consumerCnt)

	for i := range producerCnt {
		producers[i] = NewProducer(stream)
	}
	for i := range consumerCnt {
		consumers[i] = NewConsumer(stream, deadLetter, maxRetries)
	}

	return &Orchestrator{
		Stream:      stream,
		Producers:   producers,
		Consumers:   consumers,
		ProducerCnt: producerCnt,
		ConsumerCnt: consumerCnt,
		MaxRetries:  maxRetries,
		DeadLetter:  deadLetter,
		Capacity:    cap,
	}
}
