package main

/**
 * Definition for a Node.
 * type Node struct {
 *     Val int
 *     Children []*Node
 * }
 */

type Node589 struct {
	Val      int
	Children []*Node589
}

func preorder(root *Node589) []int {
	res := []int{}
	var dfs func(node *Node589)

	dfs = func(node *Node589) {
		if node == nil {
			return
		}
		res = append(res, node.Val)
		for _, child := range node.Children {
			dfs(child)
		}
	}
	dfs(root)
	return res
}
