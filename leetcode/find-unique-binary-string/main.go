package main

func findDifferentBinaryString(nums []string) string {
	n := len(nums)
	ans := make([]byte, n)
	for i := range n {
		if nums[i][i] == '0' {
			ans[i] = byte('1')
		} else {
			ans[i] = byte('0')
		}
	}
	return string(ans)
}
