/**
 * @author: baowj
 * @time: 2022/9/10 10:12:33
 */
package main

func trimBST(root *TreeNode, low int, high int) *TreeNode {
	if root == nil {
		return nil
	}
	l := trimBST(root.Left, low, high)
	r := trimBST(root.Right, low, high)
	if root.Val >= low && root.Val <= high {
		root.Left = l
		root.Right = r
		return root
	} else {
		if l == nil {
			return r
		} else {
			return l
		}
	}
}
