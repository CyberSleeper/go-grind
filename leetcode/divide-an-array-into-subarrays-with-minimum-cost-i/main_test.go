package main

import "testing"

func Test(t *testing.T) {
	tests := []struct {
		nums []int
		want int
	}{
		{[]int{1, 2, 3, 12}, 6},
		{[]int{5, 4, 3}, 12},
		{[]int{10, 3, 1, 1}, 12},
	}

	for idx, tt := range tests {
		got := minimumCost(tt.nums)
		if got != tt.want {
			t.Errorf("%d) got: %d | want: %d", idx+1, got, tt.want)
		}
	}
}
