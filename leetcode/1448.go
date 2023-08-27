package main

import "math"

func goodNodes(root *TreeNode) int {
	var dfs func(root *TreeNode, val int) int
	dfs = func(root *TreeNode, val int) int {
		if root == nil {
			return 0
		}
		res := 0
		if root.Val >= val {
			val = root.Val
			res++
		}
		res += dfs(root.Left, val)
		res += dfs(root.Right, val)
		return res
	}
	return dfs(root, math.MinInt)
}
