package main

// Just use Heap

import (
	"container/heap"
)

type IntHeap []int

func (h IntHeap) Len() int           { return len(h) }
func (h IntHeap) Less(i, j int) bool { return h[i] < h[j] }
func (h IntHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *IntHeap) Push(v any) {
	*h = append(*h, v.(int))
}
func (h *IntHeap) Pop() any {
	old := *h
	n := len(old)
	*h = old[0 : n-1]
	return old[n-1]
}

type KthLargest struct {
	K    int
	Heap *IntHeap
}

func Constructor(k int, nums []int) KthLargest {
	kthLargest := &KthLargest{
		K:    k,
		Heap: &IntHeap{},
	}
	for _, v := range nums {
		kthLargest.Add(v)
	}

	return *kthLargest
}

func (this *KthLargest) Add(val int) int {
	heap.Push(this.Heap, val)
	for this.Heap.Len() > this.K {
		heap.Pop(this.Heap)
	}
	return (*this.Heap)[0]
}

/**
 * Your KthLargest object will be instantiated and called as such:
 * obj := Constructor(k, nums);
 * param_1 := obj.Add(val);
 */
