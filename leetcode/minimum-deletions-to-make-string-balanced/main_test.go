package main

import "testing"

func Test(t *testing.T) {
	tests := []struct {
		s    string
		want int
	}{
		{
			s:    "aababbab",
			want: 2,
		},
		{
			s:    "bbaaaaabb",
			want: 2,
		},
	}

	for idx, tt := range tests {
		got := minimumDeletions(tt.s)
		if got != tt.want {
			t.Errorf("%d) got: %d | want %d", idx+1, got, tt.want)
		}
	}
}
