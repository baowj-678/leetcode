/**
 * @author: baowj
 * @time: 2023/9/6 10:43
 */
package main

func lcaDeepestLeaves(root *TreeNode) *TreeNode {
	var dfs func(node *TreeNode) (int, *TreeNode)
	dfs = func(node *TreeNode) (int, *TreeNode) {
		if node == nil {
			return 0, nil
		}
		lv, lt := dfs(node.Left)
		rv, rt := dfs(node.Right)
		if lv == rv {
			return lv + 1, node
		}
		if lv < rv {
			return rv + 1, rt
		}
		return lv + 1, lt
	}

	_, res := dfs(root)
	return res
}
