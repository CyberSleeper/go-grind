package main

func minimumDeletions(s string) int {
	n := len(s)
	prefB := make([]int, n)
	suffA := make([]int, n)

	for i := range n {
		if s[i] == 'b' {
			prefB[i] = 1
		} else {
			prefB[i] = 0
		}
		if i > 0 {
			prefB[i] += prefB[i-1]
		}
	}

	for i := n - 1; i >= 0; i-- {
		if s[i] == 'a' {
			suffA[i] = 1
		} else {
			suffA[i] = 0
		}
		if i+1 < n {
			suffA[i] += suffA[i+1]
		}
	}

	ans := min(prefB[n-1], suffA[0])
	for i := 1; i < n; i++ {
		ans = min(ans, prefB[i-1]+suffA[i])
	}

	return ans
}
