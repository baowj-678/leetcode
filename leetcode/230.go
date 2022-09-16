/**
 * @author: baowj
 * @time: 2022/8/26 21:13:18
 */
package main

func kthSmallest(root *TreeNode, k int) int {
	_, x := dfs230(root, k)
	return x
}

func dfs230(root *TreeNode, k int) (bool, int) {
	if root == nil {
		return false, 0
	}
	ok, l := dfs230(root.Left, k)
	if ok {
		return ok, l
	}
	if l == k-1 {
		return true, root.Val
	}
	ok, r := dfs230(root.Right, k-l-1)
	if ok {
		return true, r
	}
	return false, l + r + 1
}
