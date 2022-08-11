/**
 * @author: baowj
 * @time: 2022/8/5 10:14:47
 */
package main

func addOneRow(root *TreeNode, val int, depth int) *TreeNode {
	root = dfs623(root, val, depth, true)
	return root
}

func dfs623(root *TreeNode, val int, depth int, isLeft bool) *TreeNode {
	if root == nil && depth > 1 {
		return root
	}
	if depth == 1 {
		node := TreeNode{
			Val:   val,
			Left:  nil,
			Right: nil,
		}
		if isLeft {
			node.Left = root
		} else {
			node.Right = root
		}
		return &node
	}
	root.Left = dfs623(root.Left, val, depth-1, true)
	root.Right = dfs623(root.Right, val, depth-1, false)
	return root
}
