package main

import "container/list"

type LRUCache struct {
	Capacity int
	HashMap  map[int]*list.Element
	Dll      *list.List
}

type Node struct {
	Key   int
	Value int
}

func NewNode(key, value int) *Node {
	return &Node{
		Key:   key,
		Value: value,
	}
}

func Constructor(capacity int) LRUCache {
	return LRUCache{
		Capacity: capacity,
		HashMap:  make(map[int]*list.Element),
		Dll:      list.New(),
	}
}

func (c *LRUCache) Get(key int) int {
	v, exists := c.HashMap[key]
	if !exists {
		return -1
	}
	c.Dll.MoveToFront(v)
	// fmt.Println("GET", this.HashMap[key].Value, v.Value)
	el := v.Value.(*Node)
	return int(el.Value)
}

func (c *LRUCache) evict() {
	v := c.Dll.Back()
	if v == nil {
		return
	}
	delete(c.HashMap, v.Value.(*Node).Key)
	c.Dll.Remove(v)
}

func (c *LRUCache) Put(key int, value int) {
	v, exists := c.HashMap[key]
	if !exists {
		if c.Capacity == len(c.HashMap) {
			c.evict()
		}
		newNode := NewNode(key, value)
		c.HashMap[key] = c.Dll.PushFront(newNode)
		// fmt.Println("CREATE", this.HashMap[key].Value, newNode.Value)
	} else {
		c.Dll.MoveToFront(v)
		v.Value.(*Node).Value = value
		// fmt.Println("SET", this.HashMap[key].Value, el.Value)
	}
}

/**
 * Your LRUCache object will be instantiated and called as such:
 * obj := Constructor(capacity);
 * param_1 := obj.Get(key);
 * obj.Put(key,value);
 */
