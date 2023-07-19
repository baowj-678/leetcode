/**
 * @author: baowj
 * @time: 2023/7/14 10:24
 */
package main

func distributeCoins(root *TreeNode) int {
	_, _, res := dfs979(root)
	return res
}

func dfs979(root *TreeNode) (more int, less int, res int) {
	more, less, res = 0, 0, 0
	if root == nil {
		return
	}
	max := func(x, y int) int {
		if x < y {
			x = y
		}
		return x
	}
	lm, ll, lr := dfs979(root.Left)
	rm, rl, rr := dfs979(root.Right)
	res = lr + rr + max(lm, ll) + max(rm, rl)
	more = max(root.Val+lm+rm-ll-rl-1, 0)
	less = max(1+ll+rl-lm-rm-root.Val, 0)
	return
}
