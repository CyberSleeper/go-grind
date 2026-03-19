package main

import (
	"container/list"
	"fmt"
)

const debug = false

type Node struct {
	Key     int
	Value   int
	Counter int
}

type LFUCache struct {
	HashMap  map[int]*list.Element
	Dll      *list.List
	Latest   map[int]*list.Element
	Length   int
	Capacity int
}

func (this *LFUCache) debug(info string) {
	if !debug {
		return
	}

	fmt.Printf("%s = ", info)
	cur := this.Dll.Front()
	for cur != nil {
		fmt.Printf("{%d:%d:%d} ", cur.Value.(*Node).Key, cur.Value.(*Node).Value, cur.Value.(*Node).Counter)
		cur = cur.Next()
	}
	fmt.Println()
	fmt.Printf("Latest:\n")
	for k, v := range this.Latest {
		fmt.Printf("%d:%d\n", k, v.Value.(*Node).Key)
	}
	fmt.Printf("====================\n")
}

func (this *LFUCache) evict() {
	defer this.debug("EVICT")

	if this.Length == 0 {
		return
	}
	v := this.Dll.Back()
	cnt := v.Value.(*Node).Counter
	if this.Latest[cnt] == v {
		delete(this.Latest, cnt)
	}
	delete(this.HashMap, v.Value.(*Node).Key)
	this.Dll.Remove(v)
	this.Length--
}

func (this *LFUCache) Inc(el *list.Element) {
	cnt := el.Value.(*Node).Counter
	var latestRightAfter *list.Element
	if this.Latest[cnt] == el {
		latestRightAfter = el.Next()
		if latestRightAfter != nil {
			if latestRightAfter.Value.(*Node).Counter != cnt {
				delete(this.Latest, cnt)
			} else {
				this.Latest[cnt] = latestRightAfter
			}
		} else {
			delete(this.Latest, cnt)
			latestRightAfter = nil
		}
	} else {
		latestRightAfter = this.Latest[cnt]
	}

	cnt++
	el.Value.(*Node).Counter = cnt

	this.Dll.Remove(el)
	var newEl *list.Element
	if v, exists := this.Latest[cnt]; exists {
		newEl = this.Dll.InsertBefore(el.Value, v)
	} else {
		if latestRightAfter == nil {
			newEl = this.Dll.PushBack(el.Value)
		} else {
			newEl = this.Dll.InsertBefore(el.Value, latestRightAfter)
		}
	}
	this.Latest[cnt] = newEl
	this.HashMap[newEl.Value.(*Node).Key] = newEl
}

func NewNode(k, v int) *Node {
	return &Node{
		Key:     k,
		Value:   v,
		Counter: 1,
	}
}

func Constructor(capacity int) LFUCache {
	return LFUCache{
		HashMap:  make(map[int]*list.Element),
		Latest:   make(map[int]*list.Element),
		Dll:      list.New(),
		Length:   0,
		Capacity: capacity,
	}
}

func (this *LFUCache) Get(key int) int {
	defer this.debug("GET")

	v, exists := this.HashMap[key]
	if !exists {
		return -1
	}
	this.Inc(v)
	return v.Value.(*Node).Value
}

func (this *LFUCache) Put(key int, value int) {
	defer this.debug("PUT")

	v, exists := this.HashMap[key]
	if !exists {
		if this.Length == this.Capacity {
			this.evict()
		}
		this.Length++
		newNode := NewNode(key, value)
		var newEl *list.Element
		if mark, exists := this.Latest[1]; exists {
			newEl = this.Dll.InsertBefore(newNode, mark)
		} else {
			newEl = this.Dll.PushBack(newNode)
		}
		this.HashMap[key] = newEl
		this.Latest[1] = newEl
	} else {
		this.Inc(v)
		v.Value.(*Node).Value = value
	}
}

/**
 * Your LFUCache object will be instantiated and called as such:
 * obj := Constructor(capacity);
 * param_1 := obj.Get(key);
 * obj.Put(key,value);
 */
