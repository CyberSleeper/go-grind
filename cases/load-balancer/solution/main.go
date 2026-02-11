package main

import (
	"fmt"
	"hash/crc32"
	"slices"
	"sort"
	"strconv"
	"sync"
	"time"
)

const MAX_CONCURRENCY = 100

type ConsistentHash struct {
	replicas int
	keys     []uint32
	hashMap  map[uint32]string
	mu       sync.RWMutex
}

func NewConsistentHash(replicas int) *ConsistentHash {
	return &ConsistentHash{
		replicas: replicas,
		hashMap:  make(map[uint32]string),
	}
}

func (c *ConsistentHash) Add(node string) {
	c.mu.Lock()
	defer c.mu.Unlock()

	for idx := range c.replicas {
		strIdx := strconv.Itoa(idx)
		vKey := node + "#" + strIdx

		hash := crc32.ChecksumIEEE([]byte(vKey))
		c.keys = append(c.keys, hash)
		c.hashMap[hash] = node
	}

	slices.Sort(c.keys)
}

func (c *ConsistentHash) Get(key string) string {
	c.mu.RLock()
	defer c.mu.RUnlock()

	if len(c.keys) == 0 {
		return ""
	}

	hash := int(crc32.ChecksumIEEE([]byte(key)))
	idx := sort.Search(len(c.keys), func(i int) bool {
		return c.keys[i] >= uint32(hash)
	})
	if idx == len(c.keys) {
		idx = 0
	}

	return c.hashMap[c.keys[idx]]
}

func main() {
	fmt.Println("Start")
	start := time.Now()

	ring := NewConsistentHash(50)

	ring.Add("http://localhost:8081")
	ring.Add("http://localhost:8082")
	ring.Add("http://localhost:8083")

	sem := make(chan struct{}, MAX_CONCURRENCY)

	var wg sync.WaitGroup

	for i := range 100000 {
		wg.Add(1)
		sem <- struct{}{}
		go func() {
			defer wg.Done()
			defer func() { <-sem }()
			key := "user:" + strconv.Itoa(i)
			ring.Get(key)
			// fmt.Printf("Key [%s] mapped to => %s\n", key, node)

		}()
	}

	wg.Wait()

	fmt.Printf("%d conc | Time taken: %dms", MAX_CONCURRENCY, time.Since(start).Milliseconds())
}
