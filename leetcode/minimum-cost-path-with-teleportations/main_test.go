package main

import "testing"

func Test(t *testing.T) {
	tests := []struct {
		grid [][]int
		k    int
		want int
	}{
		{
			grid: [][]int{{1, 3, 3}, {2, 5, 4}, {4, 3, 5}},
			k:    2,
			want: 7,
		},
		{
			grid: [][]int{{1, 2}, {2, 3}, {3, 4}},
			k:    1,
			want: 9,
		},
		{
			grid: [][]int{{6, 7, 1, 20, 11}, {4, 5, 18, 23, 28}},
			k:    3,
			want: 46,
		},
	}

	for idx, tt := range tests {
		got := minCost(tt.grid, tt.k)
		if got != tt.want {
			t.Errorf("%d) got: %d | want %d", idx, got, tt.want)
		}
	}
}
