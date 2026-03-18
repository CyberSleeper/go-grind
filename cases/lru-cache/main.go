package main

import "container/list"

type LRUCache struct {
	Capacity int
	Size     int
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
		Size:     0,
		HashMap:  make(map[int]*list.Element),
		Dll:      list.New(),
	}
}

func (this *LRUCache) Get(key int) int {
	v, exists := this.HashMap[key]
	if !exists {
		return -1
	}
	this.Dll.MoveToFront(v)
	// fmt.Println("GET", this.HashMap[key].Value, v.Value)
	el := v.Value.(*Node)
	return int(el.Value)
}

func (this *LRUCache) evict() {
	v := this.Dll.Back()
	if v == nil {
		return
	}
	delete(this.HashMap, v.Value.(*Node).Key)
	this.Dll.Remove(v)
	this.Size--
}

func (this *LRUCache) Put(key int, value int) {
	v, exists := this.HashMap[key]
	if !exists {
		if this.Capacity == this.Size {
			this.evict()
		}
		this.Size++
		newNode := NewNode(key, value)
		this.HashMap[key] = this.Dll.PushFront(newNode)
		// fmt.Println("CREATE", this.HashMap[key].Value, newNode.Value)
	} else {
		this.Dll.MoveToFront(v)
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
