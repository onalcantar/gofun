package main

type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
}

func hasPathSum(root *TreeNode, targetSum int) bool {
    if root == nil {
		return false
	}
	
	if root.Left == nil && root.Right == nil {
		return root.Val == targetSum
	}

	newTarget := targetSum - root.Val
	return hasPathSum(root.Left, newTarget) || hasPathSum(root.Right, newTarget)
}