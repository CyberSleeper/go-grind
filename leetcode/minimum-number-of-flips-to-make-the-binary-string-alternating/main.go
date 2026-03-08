package main

func minFlips(s string) int {
	n := len(s)
	suff0 := make([]int, n+2)

	for i := n - 1; i >= 0; i-- {
		v := s[i]
		if v == '0' {
			suff0[i] = 0
		} else {
			suff0[i] = 1
		}
		if i+1 < n {
			suff0[i] += n - i - 1 - suff0[i+1]
		}
	}

	ans := min(suff0[0], n-suff0[0])
	if n&1 == 0 {
		return ans
	}

	pref0 := make([]int, n+2)

	for i, v := range s {
		if v == '0' {
			pref0[i] = 0
		} else {
			pref0[i] = 1
		}
		if i > 0 {
			pref0[i] += i - pref0[i-1]
		}
	}

	for i := 1; i < n; i++ {
		ans = min(ans, pref0[i-1]+suff0[i])
		ans = min(ans, n-pref0[i-1]-suff0[i])
	}

	return ans
}
