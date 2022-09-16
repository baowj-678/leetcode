/**
 * @author: baowj
 * @time: 2022/8/30 10:24:50
 */
package main

func insertIntoMaxTree(root *TreeNode, val int) *TreeNode {
	_, node := dfs998(root, val)
	return node
}

func dfs998(root *TreeNode, val int) (bool, *TreeNode) {
	if root == nil {
		return true, &TreeNode{Val: val, Left: nil, Right: nil}
	}
	if root.Val < val {
		node := &TreeNode{Val: val, Left: root, Right: nil}
		return true, node
	}
	ok, left := dfs998(root.Left, val)
	root.Left = left
	if !ok {
		ok, right := dfs998(root.Right, val)
		root.Right = right
		if ok {
			return true, root
		} else {
			return false, root
		}
	}
	return false, root
}
