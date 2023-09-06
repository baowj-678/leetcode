/**
 * @author: baowj
 * @time: 2023/8/31 9:58
 */
package main

func checkTree(root *TreeNode) bool {
	return root.Val == root.Left.Val+root.Right.Val
}
