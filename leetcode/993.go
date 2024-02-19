package main

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func isCousins(root *TreeNode, x int, y int) bool {
	res := false
	flag := false
	var dfs func(node *TreeNode, deep int) (int, int)
	dfs = func(node *TreeNode, deep int) (xFlag, yFlag int) {
		xFlag, yFlag = -1, -1
		if node == nil {
			return
		}
		if node.Val == x {
			xFlag = deep
		}
		if node.Val == y {
			yFlag = deep
		}

		xTmp, yTmp := dfs(node.Left, deep+1)
		if xTmp != -1 {
			xFlag = xTmp
		}
		if yTmp != -1 {
			yFlag = yTmp
		}

		xTmp, yTmp = dfs(node.Right, deep+1)
		if xTmp != -1 {
			xFlag = xTmp
		}
		if yTmp != -1 {
			yFlag = yTmp
		}
		if xFlag == yFlag && xFlag != -1 && !flag {
			if deep < xFlag-1 {
				res = true
			}
			flag = true
		}
		return
	}
	dfs(root, 0)
	return res
}
