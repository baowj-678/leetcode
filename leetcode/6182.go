/**
 * @author: baowj
 * @time: 2022/9/18 13:38:22
 */
package main

func reverseOddLevels(root *TreeNode) *TreeNode {
	reverseFunc(root.Left, root.Right, 1)
	return root
}

func reverseFunc(left *TreeNode, right *TreeNode, level int) {
	if left == nil || right == nil {
		return
	}
	if level%2 == 1 {
		left.Val, right.Val = right.Val, left.Val
	}
	reverseFunc(left.Left, right.Right, level+1)
	reverseFunc(left.Right, right.Left, level+1)
}
