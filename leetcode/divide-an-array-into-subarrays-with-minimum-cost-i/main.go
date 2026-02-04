package main

const INF = 1e9 + 7

func minimumCost(nums []int) int {
	ans := make([]int, 2)
	for i := range ans {
		ans[i] = INF
	}

	for _, v := range nums[1:] {
		if v < ans[1] {
			ans[1] = v
		}
		if ans[1] < ans[0] {
			ans[0], ans[1] = ans[1], ans[0]
		}
	}

	return ans[0] + ans[1] + nums[0]
}
