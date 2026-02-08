package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

const INF = int(1e9) + 7

func AbsInt(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func depth(root *TreeNode) (int, bool) {
	if root == nil {
		return 0, true
	}

	depLeft, validLeft := depth(root.Left)
	depRight, validRight := depth(root.Right)

	return max(depLeft, depRight) + 1, validLeft && validRight && AbsInt(depLeft-depRight) < 2
}

func isBalanced(root *TreeNode) bool {
	_, valid := depth(root)
	return valid
}
