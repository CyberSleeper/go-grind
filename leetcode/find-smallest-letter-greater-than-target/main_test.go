package main

import "testing"

func Test(t *testing.T) {
	tests := []struct {
		letters []byte
		target  byte
		want    byte
	}{
		{[]byte{'c', 'f', 'j'}, 'a', 'c'},
		{[]byte{'c', 'f', 'j'}, 'c', 'f'},
		{[]byte{'c', 'f', 'j'}, 'd', 'f'},
		{[]byte{'c', 'f', 'j'}, 'z', 'c'},
		{[]byte{'x', 'x', 'y', 'y'}, 'z', 'x'},
	}

	for idx, tt := range tests {
		got := nextGreatestLetter(tt.letters, tt.target)
		if got != tt.want {
			t.Errorf("%d) got: %c | want: %c", idx+1, got, tt.want)
		}
	}
}
