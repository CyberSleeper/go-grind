package main

import "testing"

func Test(t *testing.T) {
	tests := []struct {
		root *TreeNode
		want bool
	}{
		{
			root: &TreeNode{
				Val: 3,
				Left: &TreeNode{
					Val: 9,
				},
				Right: &TreeNode{
					Val: 20,
					Left: &TreeNode{
						Val: 15,
					},
					Right: &TreeNode{
						Val: 7,
					},
				},
			},
			want: true,
		},
	}

	for idx, tt := range tests {
		got := isBalanced(tt.root)
		if got != tt.want {
			t.Errorf("%d) got: %v | want: %v", idx+1, got, tt.want)
		}
	}
}
