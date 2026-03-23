package main

import "slices"

const maxData = 30007

type DSU struct {
	Parent []int
}

func NewDSU(n int) *DSU {
	parent := make([]int, n+5)

	for i := range n + 1 {
		parent[i] = i
	}

	return &DSU{
		Parent: parent,
	}
}

func (d *DSU) Join(u, v int) bool {
	if u < 0 || v < 0 {
		return false
	}

	u = d.FindParent(u)
	v = d.FindParent(v)

	if u == v {
		return false
	}

	d.Parent[u] = v
	return true
}

func (d *DSU) FindParent(x int) int {
	if d.Parent[x] == x {
		return x
	}
	d.Parent[x] = d.FindParent(d.Parent[x])
	return d.Parent[x]
}

type SummaryRanges struct {
	Left          *DSU
	Right         *DSU
	ActiveDataSet map[int]struct{}
}

func Constructor() SummaryRanges {
	return SummaryRanges{
		Left:          NewDSU(maxData),
		Right:         NewDSU(maxData),
		ActiveDataSet: make(map[int]struct{}),
	}
}

func (c *SummaryRanges) AddNum(value int) {
	// add 1 so 0 still can access -1
	value++

	c.Left.Join(value, value-1)
	c.Right.Join(value, value+1)
	c.ActiveDataSet[value] = struct{}{}
}

func (c *SummaryRanges) GetIntervals() [][]int {
	activeDataList := make([][]int, 0, len(c.ActiveDataSet))
	newActiveDataSet := make(map[int]struct{})

	for k := range c.ActiveDataSet {
		leftMost := c.Left.FindParent(k) + 1
		rightMost := c.Right.FindParent(k) - 1

		if k == leftMost {
			activeDataList = append(activeDataList, []int{leftMost - 1, rightMost - 1})
			newActiveDataSet[leftMost] = struct{}{}
		}
	}

	slices.SortFunc(activeDataList, func(a, b []int) int {
		return a[0] - b[0]
	})
	c.ActiveDataSet = newActiveDataSet

	return activeDataList
}

/**
 * Your SummaryRanges object will be instantiated and called as such:
 * obj := Constructor();
 * obj.AddNum(value);
 * param_2 := obj.GetIntervals();
 */
