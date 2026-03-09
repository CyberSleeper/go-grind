package main

import (
	"slices"
)

const arrSize = 4
const inf = 1_000_000_007

func maximumTotalDamage(power []int) int64 {
	cur := make([]int64, arrSize+3)
	prev := make([]int64, arrSize+3)

	slices.Sort(power)

	bef := -inf

	for _, v := range power {
		prev, cur = cur, prev
		for i := range 4 {
			cur[i] = 0
		}

		diff := v - bef
		bef = v

		for j := 3; j >= 0; j-- {
			target := min(3, j+diff)
			cur[target] = max(cur[target], prev[j])
		}

		cur[0] = max(cur[0], cur[3])
		cur[0] += int64(v)
	}

	ans := int64(0)
	for _, v := range cur {
		ans = max(ans, v)
	}
	return ans
}
