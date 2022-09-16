/**
 * @author: baowj
 * @time: 2022/8/25 16:38:04
 */
package main

import "strconv"

func printTree(root *TreeNode) [][]string {
	h := height(root) - 1
	list := make([][]string, h+1)
	n := pow2(h+1) - 1
	for i := 0; i < len(list); i++ {
		list[i] = make([]string, n)
	}
	dfs655(h, 0, n/2, root, list)
	return list
}

func dfs655(h int, r int, c int, root *TreeNode, list [][]string) {
	if root == nil {
		return
	}
	list[r][c] = strconv.Itoa(root.Val)
	delta := pow2(h - r - 1)
	dfs655(h, r+1, c-delta, root.Left, list)
	dfs655(h, r+1, c+delta, root.Right, list)
}

func height(root *TreeNode) int {
	if root == nil {
		return 0
	}
	l := height(root.Left)
	r := height(root.Right)
	if l > r {
		return l + 1
	} else {
		return r + 1
	}
}

func pow2(p int) int {
	res := 1
	for i := 0; i < p; i++ {
		res *= 2
	}
	return res
}
