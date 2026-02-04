package main

import "testing"

func Test(t *testing.T) {
	tests := []struct {
		nums []int
		k    int
		dist int
		want int64
	}{
		{[]int{1, 3, 2, 6, 4, 2}, 3, 3, 5},
		{[]int{10, 1, 2, 2, 2, 1}, 4, 3, 15},
		{[]int{10, 8, 18, 9}, 3, 1, 36},
		{[]int{2, 5, 3, 5, 7, 4, 3}, 3, 3, 9},
	}

	for idx, tt := range tests {
		got := minimumCost(tt.nums, tt.k, tt.dist)
		if got != tt.want {
			t.Errorf("%d) got: %d | want: %d", idx+1, got, tt.want)
		}
	}
}
