## The Load Balancer (Consistent Hashing)

### 📌 Concept

In distributed systems, we need to decide which server stores a specific key (Sharding). The naive approach uses `hash(key) % NumberOfNodes`.

**The Scenario:** You have 3 nodes (`A`, `B`, `C`).
* If you add a 4th node (`D`), the formula changes to `hash % 4`.
* **The Disaster:** This causes **~75% of all keys to move** to different servers instantly. Your database cache vanishes, and the network floods with rebalancing traffic.

**Consistent Hashing** maps both Nodes and Keys onto a 32-bit "Ring" (a sorted circle). When a node is added, it only takes a small slice of keys from its neighbor, leaving the rest untouched.

### 📝 Task

Implement a `ConsistentHash` struct that supports **Virtual Nodes** (replicas) to ensure balanced distribution.

1.  **Add(node):** Loop `replicas` times. Create virtual keys (e.g., `NodeA#1`, `NodeA#2`...). Hash them and insert them into a sorted slice called `keys`.
2.  **Get(key):**
    * Hash the incoming key.
    * Search the sorted `keys` slice for the first value that is **>=** the key's hash (Moving Clockwise).
    * **Wrap Around:** If the key is larger than all points on the ring, it belongs to the first node (Index 0).
3.  **Constraint:** You must use `crc32.ChecksumIEEE` for hashing and `sort.Search` for looking up the position.

### 🚫 Starter Code

```go
package main

import (
	"fmt"
	"hash/crc32"
	"sort"
	"strconv"
)

type ConsistentHash struct {
	replicas int            // Number of virtual nodes per real node
	keys     []int          // Sorted list of hash points on the ring
	hashMap  map[int]string // Map: Hash Point -> Real Node Name
}

func New(replicas int) *ConsistentHash {
	return &ConsistentHash{
		replicas: replicas,
		hashMap:  make(map[int]string),
	}
}

func (c *ConsistentHash) Add(node string) {
	// TODO:
	// 1. Loop 'c.replicas' times.
	// 2. Create a virtual key: strconv.Itoa(i) + node
	// 3. Hash it using crc32.ChecksumIEEE([]byte(virtualKey))
	// 4. Store the hash in c.keys and map it to the real node name in c.hashMap
	// 5. SORT the c.keys slice! (Crucial)
}

func (c *ConsistentHash) Get(key string) string {
	if len(c.keys) == 0 {
		return ""
	}

	// TODO:
	// 1. Hash the incoming key (crc32)
	// 2. Use sort.Search to find the index of the first node hash >= key hash
	// 3. Handle the "Wrap Around" case: if index == len(c.keys), reset index to 0.
	
	return "" // Return the real node name from c.hashMap
}

func main() {
	// Create a ring with 3 replicas per node
	ring := New(3)
	
	// Add 3 physical nodes
	ring.Add("Node A")
	ring.Add("Node B")
	ring.Add("Node C")

	// Test a key
	key := "user:12345"
	node := ring.Get(key)
	
	fmt.Printf("Key [%s] mapped to => %s\n", key, node)
}
```

### 💡 Hint

`sort.Search(n, f)` returns the smallest index `i` where `f(i)` is true. If it can't find one, it returns `n`. In the context of a Ring, returning `n` means we went past the end of the line—so the "next" node is technically the start of the array (`index 0`).