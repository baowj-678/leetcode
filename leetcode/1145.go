/**
 * @author: baowj
 * @time: 2023/2/3 10:12:36
 * @email: bwj_678@qq.com
 */
package main

func btreeGameWinningMove(root *TreeNode, n int, x int) bool {
	_, left, right := dfs1145(root, x)
	if left > n-left || right > n-right || n-(left+right+1) > left+right+1 {
		return true
	}
	return false
}

func dfs1145(root *TreeNode, x int) (ok bool, left int, right int) {
	left, right = 0, 0
	ok = false
	if root == nil {
		return
	}
	ok, left, right = dfs1145(root.Left, x)
	if ok {
		return
	}
	leftSum := left + right
	ok, left, right = dfs1145(root.Right, x)
	if ok {
		return
	}
	rightSum := left + right
	left = leftSum
	if root.Left != nil {
		left++
	}
	right = rightSum
	if root.Right != nil {
		right++
	}
	if root.Val == x {
		ok = true
	}
	return
}
