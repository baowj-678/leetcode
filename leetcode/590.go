package main

func postorder(root *Node589) []int {
	res := []int{}
	var dfs func(node *Node589)
	dfs = func(node *Node589) {
		if node == nil {
			return
		}
		for _, child := range node.Children {
			dfs(child)
		}
		res = append(res, node.Val)
	}
	dfs(root)
	return res
}
