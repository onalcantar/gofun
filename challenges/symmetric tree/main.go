package main

type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
}

func isSymmetric(root *TreeNode) bool {
	if root == nil {
		return true
	}

	return isMirror(root.Left, root.Right)
}

func isMirror(leftNode, rightNode *TreeNode) bool {
	if leftNode == nil && rightNode == nil {
		return true
	}

	if leftNode == nil || rightNode == nil || leftNode.Val != rightNode.Val  {
		return false
	}

	return isMirror(leftNode.Left, rightNode.Right) && isMirror(leftNode.Right, rightNode.Left)
}