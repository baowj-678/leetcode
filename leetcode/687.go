/**
 * @author: baowj
 * @time: 2022/9/5 10:35:47
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
func longestUnivaluePath(root *TreeNode) int {
	m, _, _ := dfs687(root)
	if m == 0 {
		return 0
	}
	return m - 1
}

func dfs687(root *TreeNode) (maxNode int, maxValNode int, val int) {
	if root == nil {
		maxNode = 0
		maxValNode = 0
		val = -100000
		return
	}
	l, ln, lv := dfs687(root.Left)
	r, rn, rv := dfs687(root.Right)
	maxNode = 0
	maxValNode = 1
	val = root.Val
	if l > maxNode {
		maxNode = l
	}
	if r > maxNode {
		maxNode = r
	}
	if lv == val {
		if maxValNode < 1+ln {
			maxValNode = 1 + ln
		}
	}
	if rv == val {
		if maxValNode < 1+rn {
			maxValNode = 1 + rn
		}
	}
	if rv == val && lv == val {
		if maxNode < 1+ln+rn {
			maxNode = 1 + ln + rn
		}
	}
	if maxValNode > maxNode {
		maxNode = maxValNode
	}
	return
}
