package main

// Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isBalanced(root *TreeNode) bool {
	if root == nil {
		return true
	}
	h := height(root.Left) - height(root.Right)
	return h >= -1 && h <= 1 && isBalanced(root.Left) && isBalanced(root.Right)
}

func height(node *TreeNode) int {
	if node == nil {
		return 0
	}
	l := height(node.Left) + 1
	r := height(node.Right) + 1
	if l > r {
		return l
	}
	return r
}
