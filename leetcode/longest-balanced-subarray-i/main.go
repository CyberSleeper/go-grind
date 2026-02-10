package main

func longestBalanced(nums []int) int {
	cnt := make([]int, 100007)
	n := len(nums)
	par := make([]int, 2)
	ans := 0

	for le := 0; le < n; le++ {
		for ri := le; ri < n; ri++ {
			if cnt[nums[ri]] == 0 {
				par[nums[ri]%2]++
			}
			cnt[nums[ri]]++
			if par[0] == par[1] {
				ans = max(ans, ri-le+1)
			}
		}
		for ri := le; ri < n; ri++ {
			cnt[nums[ri]] = 0
		}
		par[0], par[1] = 0, 0
	}
	return ans
}
