package main

import (
	"container/heap"
)

const INFF = int64(1e18) + 7

type Item struct {
	Val int
	Idx int
}

// --- Priority Queue Boilerplate (Copy-Paste this) ---
type PriorityQueue []Item

func (pq PriorityQueue) Len() int            { return len(pq) }
func (pq PriorityQueue) Less(i, j int) bool  { return pq[i].Val > pq[j].Val }
func (pq PriorityQueue) Swap(i, j int)       { pq[i], pq[j] = pq[j], pq[i] }
func (pq *PriorityQueue) Push(x interface{}) { *pq = append(*pq, x.(Item)) }
func (pq *PriorityQueue) Pop() interface{} {
	old := *pq
	n := len(old)
	item := old[n-1]
	*pq = old[0 : n-1]
	return item
}

// ---------------------------------------------------

func minimumCost(nums []int, k int, dist int) int64 {
	n := len(nums)
	pq := &PriorityQueue{}
	counted := make([]bool, n)
	heap.Init(pq)
	ans := INFF
	var currSum int64 = int64(nums[0])
	le := 1
	sz := 0
	k--

	for ri := 1; ri < n; ri++ {
		currSum += int64(nums[ri])
		counted[ri] = true
		heap.Push(pq, Item{nums[ri], ri})
		sz++

		if ri-le > dist {
			if counted[le] {
				counted[le] = false
				currSum -= int64(nums[le])
				sz--
			}
			le++
		}

		for !counted[(*pq)[0].Idx] || sz > k {
			if counted[(*pq)[0].Idx] {
				counted[(*pq)[0].Idx] = false
				currSum -= int64((*pq)[0].Val)
				sz--
			}
			heap.Pop(pq)
		}
		if sz == k {
			ans = min(ans, currSum)
		}
	}
	return ans
}
