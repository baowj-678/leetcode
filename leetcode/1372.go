/**
 * @author: baowj
 * @time: 2022/9/5 11:32:23
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
func longestZigZag(root *TreeNode) int {
	l, _ := dfs1372(root, true)
	r, _ := dfs1372(root, false)
	if l < r {
		return r
	} else {
		return l
	}
}

func dfs1372(root *TreeNode, isLeft bool) (max int, nodeMax int) {
	if root == nil {
		max = 0
		nodeMax = 0
		return
	}
	l, ln := dfs1372(root.Left, true)
	r, rn := dfs1372(root.Right, false)
	max = l
	if r > max {
		max = r
	}

	if ln+1 > max {
		max = ln + 1
	}
	if rn+1 > max {
		max = rn + 1
	}

	if isLeft {
		nodeMax = rn + 1
	} else {
		nodeMax = ln + 1
	}
	if max < nodeMax {
		max = nodeMax
	}
	return
}
