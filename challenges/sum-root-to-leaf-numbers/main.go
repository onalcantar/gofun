package main

import (
	"strconv"
	"slices"
)

type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
}

func sumNumbers(root *TreeNode) int {
	var dfs func(node *TreeNode, path []byte) int

	dfs = func(node *TreeNode, path []byte) int {
		if node == nil {
			return 0
		}

		path = append(path, byte('0' + node.Val))

		if node.Left == nil && node.Right == nil {
			num, _ := strconv.Atoi(string(path))
			return num
		}

		left := dfs(node.Left, slices.Clone(path))
		right := dfs(node.Right, slices.Clone(path))
		return left + right
	}

	return dfs(root, []byte{})
}