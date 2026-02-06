package main

import "testing"

func Test(t *testing.T) {
	tests := []struct {
		nums []int
		k    int
		want int
	}{
		{[]int{2, 1, 5}, 2, 1},
		{[]int{1, 6, 2, 9}, 3, 2},
		{[]int{4, 6}, 2, 0},
	}

	for idx, tt := range tests {
		got := minRemoval(tt.nums, tt.k)
		if got != tt.want {
			t.Errorf("%d) got %d | want: %d", idx+1, got, tt.want)
		}
	}
}
