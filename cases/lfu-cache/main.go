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

func (c *LFUCache) debug(info string) {
	if !debug {
		return
	}

	fmt.Printf("%s = ", info)
	cur := c.Dll.Front()
	for cur != nil {
		fmt.Printf("{%d:%d:%d} ", cur.Value.(*Node).Key, cur.Value.(*Node).Value, cur.Value.(*Node).Counter)
		cur = cur.Next()
	}
	fmt.Println()
	fmt.Printf("Latest:\n")
	for k, v := range c.Latest {
		fmt.Printf("%d:%d\n", k, v.Value.(*Node).Key)
	}
	fmt.Printf("====================\n")
}

func (c *LFUCache) evict() {
	defer c.debug("EVICT")

	if c.Length == 0 {
		return
	}
	v := c.Dll.Back()
	cnt := v.Value.(*Node).Counter
	if c.Latest[cnt] == v {
		delete(c.Latest, cnt)
	}
	delete(c.HashMap, v.Value.(*Node).Key)
	c.Dll.Remove(v)
	c.Length--
}

func (c *LFUCache) Inc(el *list.Element) *list.Element {
	cnt := el.Value.(*Node).Counter
	var latestRightAfter *list.Element
	if c.Latest[cnt] == el {
		latestRightAfter = el.Next()
		if latestRightAfter != nil {
			if latestRightAfter.Value.(*Node).Counter != cnt {
				delete(c.Latest, cnt)
			} else {
				c.Latest[cnt] = latestRightAfter
			}
		} else {
			delete(c.Latest, cnt)
			latestRightAfter = nil
		}
	} else {
		latestRightAfter = c.Latest[cnt]
	}

	cnt++
	el.Value.(*Node).Counter = cnt

	c.Dll.Remove(el)
	var newEl *list.Element
	if v, exists := c.Latest[cnt]; exists {
		newEl = c.Dll.InsertBefore(el.Value, v)
	} else {
		if latestRightAfter == nil {
			newEl = c.Dll.PushBack(el.Value)
		} else {
			newEl = c.Dll.InsertBefore(el.Value, latestRightAfter)
		}
	}
	c.Latest[cnt] = newEl
	c.HashMap[newEl.Value.(*Node).Key] = newEl

	return newEl
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

func (c *LFUCache) Get(key int) int {
	defer c.debug("GET")

	v, exists := c.HashMap[key]
	if !exists {
		return -1
	}
	newEl := c.Inc(v)
	return newEl.Value.(*Node).Value
}

func (c *LFUCache) Put(key int, value int) {
	defer c.debug("PUT")

	if c.Capacity == 0 {
		return
	}
	v, exists := c.HashMap[key]
	if !exists {
		if c.Length == c.Capacity {
			c.evict()
		}
		c.Length++
		newNode := NewNode(key, value)
		var newEl *list.Element
		if mark, exists := c.Latest[1]; exists {
			newEl = c.Dll.InsertBefore(newNode, mark)
		} else {
			newEl = c.Dll.PushBack(newNode)
		}
		c.HashMap[key] = newEl
		c.Latest[1] = newEl
	} else {
		newEl := c.Inc(v)
		newEl.Value.(*Node).Value = value
	}
}

/**
 * Your LFUCache object will be instantiated and called as such:
 * obj := Constructor(capacity);
 * param_1 := obj.Get(key);
 * obj.Put(key,value);
 */
