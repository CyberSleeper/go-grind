package main

import "testing"

func Test(t *testing.T) {
	tests := []struct {
		source   string
		target   string
		original []string
		changed  []string
		cost     []int
		want     int64
	}{
		{
			source:   "abcd",
			target:   "acbe",
			original: []string{"a", "b", "c", "c", "e", "d"},
			changed:  []string{"b", "c", "b", "e", "b", "e"},
			cost:     []int{2, 5, 5, 1, 2, 20},
			want:     28,
		},
		{
			source:   "abcdefgh",
			target:   "acdeeghh",
			original: []string{"bcd", "fgh", "thh"},
			changed:  []string{"cde", "thh", "ghh"},
			cost:     []int{1, 3, 5},
			want:     9,
		},
		{
			source:   "abcdefgh",
			target:   "addddddd",
			original: []string{"bcd", "defgh"},
			changed:  []string{"ddd", "ddddd"},
			cost:     []int{100, 1578},
			want:     -1,
		},
	}

	for idx, tt := range tests {
		got := minimumCost(tt.source, tt.target, tt.original, tt.changed, tt.cost)
		if got != tt.want {
			t.Errorf("%d) got: %d | want: %d", idx+1, got, tt.want)
		}
	}
}
