package main

import "sort"

const INFF = 1e18 + 7

type Graph struct {
	Dist [][]int64
}

func NewGraph(sz int) *Graph {
	g := &Graph{Dist: make([][]int64, sz)}
	for i := 0; i < sz; i++ {
		g.Dist[i] = make([]int64, sz)
		for j := 0; j < sz; j++ {
			g.Dist[i][j] = INFF
		}
		g.Dist[i][i] = 0
	}
	return g
}

func (g *Graph) AddEdge(u, v, w int) {
	g.Dist[u][v] = min(g.Dist[u][v], int64(w))
}

// TODO: Implement this function
func minimumCost(source string, target string, original []string, changed []string, cost []int) int64 {
	N := len(source)
	M := len(original)

	changedLenSet := make(map[int]struct{})
	for _, v := range original {
		curr := len(v)
		changedLenSet[curr] = struct{}{}
	}
	changedLenList := make([]int, 0, len(changedLenSet))
	for key := range changedLenSet {
		changedLenList = append(changedLenList, key)
	}
	sort.Ints(changedLenList)

	strTranslator := make(map[string]int)

	graph := NewGraph(2*M + 7)
	var oriTranslated []int

	cnt := 1
	for _, v := range original {
		cur := -1
		if _, exists := strTranslator[v]; !exists {
			strTranslator[v] = cnt
			cur = cnt
			cnt++
		} else {
			cur = strTranslator[v]
		}
		oriTranslated = append(oriTranslated, cur)
	}

	for i, v := range changed {
		cur := -1
		if _, exists := strTranslator[v]; !exists {
			strTranslator[v] = cnt
			cur = cnt
			cnt++
		} else {
			cur = strTranslator[v]
		}
		graph.AddEdge(oriTranslated[i], cur, cost[i])
	}

	for k := 1; k <= cnt; k++ {
		for i := 1; i <= cnt; i++ {
			for j := 1; j <= cnt; j++ {
				graph.Dist[i][j] = min(graph.Dist[i][j], graph.Dist[i][k]+graph.Dist[k][j])
			}
		}
	}

	dp := make([]int64, N+7)

	for i := 0; i < N; i++ {
		dp[i] = INFF
		if source[i] == target[i] {
			var prev int64
			if i > 0 {
				prev = dp[i-1]
			} else {
				prev = 0
			}
			dp[i] = min(dp[i], prev)
		}
		for _, diff := range changedLenList {
			bef := i - diff + 1
			if bef < 0 {
				break
			}
			currSource := source[bef : i+1]
			currTarget := target[bef : i+1]

			if _, exists := strTranslator[currSource]; !exists {
				continue
			}

			var cost int64
			cost = 0
			if bef > 0 {
				cost = dp[bef-1]
			}
			sourceInt := strTranslator[currSource]
			targetInt := strTranslator[currTarget]
			dp[i] = min(dp[i], cost+graph.Dist[sourceInt][targetInt])
		}
	}

	if dp[N-1] == INFF {
		return -1
	}

	return dp[N-1]
}
