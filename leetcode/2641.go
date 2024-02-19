package main

import "fmt"

func replaceValueInTree(root *TreeNode) *TreeNode {
	deepSum := []int{}
	var dfs func(node *TreeNode, deep int) int

	dfs = func(node *TreeNode, deep int) int {
		if node == nil {
			return 0
		}
		if len(deepSum) <= deep {
			deepSum = append(deepSum, 0)
		}
		deepSum[deep] += node.Val
		v := node.Val
		l := dfs(node.Left, deep+1)
		r := dfs(node.Right, deep+1)

		node.Val = l + r
		return v
	}
	dfs(root, 0)
	fmt.Println(deepSum)

	var dfsf func(node *TreeNode, sonsSum int, deep int)
	dfsf = func(node *TreeNode, sonsSum int, deep int) {
		if node == nil {
			return
		}
		v := node.Val
		node.Val = deepSum[deep] - sonsSum
		dfsf(node.Left, v, deep+1)
		dfsf(node.Right, v, deep+1)
	}
	dfsf(root, deepSum[0], 0)
	return root
}
