/**
 * @author: baowj
 * @time: 2023/4/18 10:36
 */
package main

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func maxAncestorDiff(root *TreeNode) int {
	return dfs1026(root.Val, root.Val, root)
}

func dfs1026(maxV, minV int, root *TreeNode) int {
	if root == nil {
		return 0
	}
	abs := func(a int) int {
		if a < 0 {
			a = -a
		}
		return a
	}

	max := func(a, b int) int {
		if a < b {
			a = b
		}
		return a
	}

	min := func(a, b int) int {
		if a > b {
			a = b
		}
		return a
	}

	res := 0
	res = max(res, abs(maxV-root.Val))
	res = max(res, abs(minV-root.Val))

	maxV = max(maxV, root.Val)
	minV = min(minV, root.Val)

	res = max(res, dfs1026(maxV, minV, root.Left))
	res = max(res, dfs1026(maxV, minV, root.Right))

	return res
}
