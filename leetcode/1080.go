/**
 * @author: baowj
 * @time: 2023/5/22 14:03
 */
package main

import "math"

func sufficientSubset(root *TreeNode, limit int) *TreeNode {
	root, s := dfs1080(root, limit, 0)
	if s < limit {
		root = nil
	}
	return root
}

func dfs1080(root *TreeNode, limit, sum int) (*TreeNode, int) {
	sum += root.Val
	if root.Left == nil && root.Right == nil {
		return root, sum
	}
	var tmp, max int = 0, math.MinInt
	if root.Left != nil {
		root.Left, tmp = dfs1080(root.Left, limit, sum)
		if tmp < limit {
			root.Left = nil
		}
		if tmp > max {
			max = tmp
		}
	}
	if root.Right != nil {
		root.Right, tmp = dfs1080(root.Right, limit, sum)
		if tmp < limit {
			root.Right = nil
		}
		if tmp > max {
			max = tmp
		}
	}
	return root, max
}
