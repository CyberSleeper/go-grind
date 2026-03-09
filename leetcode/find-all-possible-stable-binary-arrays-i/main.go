package main

const MOD = 1_000_000_007

func numberOfStableArrays(zero int, one int, limit int) int {
	DP0 := make([][]int, 207)
	DP1 := make([][]int, 207)
	for i := range 202 {
		DP0[i] = make([]int, 207)
		DP1[i] = make([]int, 207)
	}

	DP0[0][0] = 1
	DP1[0][0] = 1

	for sz := 1; sz <= zero+one; sz++ {
		for i := 0; i <= min(sz, zero); i++ {
			j := sz - i
			if j > one {
				continue
			}
			for k := 1; k <= limit && i-k >= 0; k++ {
				DP0[i][j] += DP1[i-k][j]
			}
			for k := 1; k <= limit && j-k >= 0; k++ {
				DP1[i][j] += DP0[i][j-k]
			}
			DP0[i][j] %= MOD
			DP1[i][j] %= MOD
		}
	}

	return (DP0[zero][one] + DP1[zero][one]) % MOD
}
