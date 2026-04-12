package main

import (
	"fmt"
	"strconv"
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
	processedDataCnt int
}

type DataStream struct {
	Stream chan DataEntry
	cap    int
}

type Orchestrator struct {
	Producers   []*Producer
	Consumers   []*Consumer
	Stream      *DataStream
	Capacity    int
	ProducerCnt int
	ConsumerCnt int
}

type DataEntry struct {
	value      string
	producerId int
	consumerId int
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

func NewConsumer(d *DataStream) *Consumer {
	return &Consumer{
		Stream: d,
	}
}

func NewOrchestrator(cap, producerCnt, consumerCnt int) *Orchestrator {
	stream := NewDataStream(cap)
	producers := make([]*Producer, producerCnt)
	consumers := make([]*Consumer, consumerCnt)

	for i := range producerCnt {
		producers[i] = NewProducer(stream)
	}
	for i := range consumerCnt {
		consumers[i] = NewConsumer(stream)
	}

	return &Orchestrator{
		Stream:      stream,
		Producers:   producers,
		Consumers:   consumers,
		ProducerCnt: producerCnt,
		ConsumerCnt: consumerCnt,
	}
}

func main() {
	orc := NewOrchestrator(5, 5, 5)
	var input chan DataEntry
	var output chan DataEntry
	var terminate chan int

	input = make(chan DataEntry, 10)
	output = make(chan DataEntry, 10)
	terminate = make(chan int, 1)

	for i := range orc.ProducerCnt {
		go func() {
			for {
				data := <-input
				data.producerId = i
				orc.Producers[i].Enqueue(data)
			}
		}()
	}

	for i := range orc.ConsumerCnt {
		go func() {
			for {
				val, ok := orc.Consumers[i].Dequeue()
				if ok {
					val.consumerId = i
					output <- val
				} else {
					break
				}
			}
		}()
	}

	for i := range 10000 {
		go func() {
			curData := DataEntry{
				value:      strconv.Itoa(i),
				producerId: -1,
				consumerId: -1,
			}

			input <- curData
		}()
	}

	go func() {
	Mainloop:
		for {
			select {
			case <-output:
				continue
			case <-terminate:
				break Mainloop
			}
		}
	}()

	for {
		var in string
		fmt.Print("Enter your message: ")
		fmt.Scanln(&in)
		if in == "exit" {
			for i, v := range orc.Producers {
				fmt.Println("Producer", i, v.processedDataCnt)
			}
			for i, v := range orc.Consumers {
				fmt.Println("Consumer", i, v.processedDataCnt)
			}
			break
		}
		d := DataEntry{
			value: in,
		}
		input <- d
	}
}
