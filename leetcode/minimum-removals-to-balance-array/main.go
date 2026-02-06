package main

import "slices"

func minRemoval(nums []int, k int) int {
	n := len(nums)
	slices.Sort(nums)
	ptr1 := 0
	ans := n - 1
	for ptr2 := range n {
		for nums[ptr1]*k < nums[ptr2] {
			ptr1++
		}
		ans = min(ans, n-ptr2+ptr1-1)
	}

	return ans
}
