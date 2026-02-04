package main

import (
	"slices"
)

const INF = (int)(1e9 + 7)

type Sweep struct {
	X    int
	Y    int
	Cost int
	Val  int
}

func minCost(grid [][]int, k int) int {
	N := len(grid)
	M := len(grid[0])

	best := make([][][]int, N+3)
	sortedGrid := make([]Sweep, 0, N*M)

	for i := range N + 2 {
		best[i] = make([][]int, M+3)
		for j := range M + 2 {
			best[i][j] = make([]int, k+3)
			for ii := range k + 2 {
				best[i][j][ii] = INF
			}
		}
	}

	for t := range k + 1 {
		best[0][0][t] = 0
		smol := make([]int, 10007)
		if t > 0 {
			sortedGrid = sortedGrid[:0]
			for i := range N {
				for j := range M {
					sortedGrid = append(sortedGrid, Sweep{i, j, best[i][j][t-1], grid[i][j]})
				}
			}
			slices.SortFunc(sortedGrid, func(a, b Sweep) int {
				return b.Val - a.Val
			})
			smol[sortedGrid[0].Val] = sortedGrid[0].Cost
			for i := 1; i < len(sortedGrid); i++ {
				sortedGrid[i].Cost = min(sortedGrid[i].Cost, sortedGrid[i-1].Cost)
				smol[sortedGrid[i].Val] = sortedGrid[i].Cost
			}
		}

		ptr := 0

		for i := range N {
			for j := range M {
				for ptr < len(sortedGrid) && sortedGrid[ptr].Val >= grid[i][j] {
					ptr++
				}

				curr := best[i][j][t]
				if i > 0 {
					curr = min(curr, best[i-1][j][t]+grid[i][j])
				}
				if j > 0 {
					curr = min(curr, best[i][j-1][t]+grid[i][j])
				}
				if t > 0 {
					curr = min(curr, smol[grid[i][j]])
				}

				best[i][j][t] = curr
			}
		}
	}

	ans := INF
	for i := range k + 1 {
		ans = min(ans, best[N-1][M-1][i])
	}

	return ans
}
