/**
 * @author: baowj
 * @time: 2023/2/6 20:24:34
 * @email: bwj_678@qq.com
 */
package main

func evaluateTree(root *TreeNode) bool {
	return dfs2331(root)
}

func dfs2331(root *TreeNode) bool {
	if root == nil {
		return false
	}
	if root.Val == 0 {
		return false
	} else if root.Val == 1 {
		return true
	}
	left := dfs2331(root.Left)
	right := dfs2331(root.Right)
	if root.Val == 2 {
		return left || right
	} else {
		return left && right
	}
}
