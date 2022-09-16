/**
 * @author: baowj
 * @time: 2022/8/21 17:43:45
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
func amountOfTime(root *TreeNode, start int) int {
	if root.Val == start {
		_, l := dfs6154(root.Left, start)
		_, r := dfs6154(root.Right, start)
		if l > r {
			return l
		} else {
			return r
		}
	} else {
		ml, l := dfs6154(root.Left, start)
		mr, r := dfs6154(root.Right, start)
		if ml == 0 {
			if mr > l+r+1 {
				return mr
			} else {
				return l + r + 1
			}
		} else {
			if ml > l+r+1 {
				return ml
			} else {
				return l + r + 1
			}
		}
	}
}

func dfs6154(root *TreeNode, start int) (maxDis int, dis int) {
	if root == nil {
		return 0, -1
	}
	maxDis, dis = 0, 0
	x, y := dfs6154(root.Left, start)
	if x > maxDis {
		maxDis = x
	}
	if y > dis {
		dis = y
	}
	x, y = dfs6154(root.Right, start)
	if x > maxDis {
		maxDis = x
	}
	if y > dis {
		dis = y
	}
	dis++
	if root.Val == start {
		maxDis = dis
		dis = 0
	}
	return maxDis, dis
}
