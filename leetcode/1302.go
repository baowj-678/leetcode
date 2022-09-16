/**
 * @author: baowj
 * @time: 2022/8/17 10:04:04
 */
package main

func deepestLeavesSum(root *TreeNode) int {
	sum, maxLevel := 0, 0
	dfs1302(root, 0, &maxLevel, &sum)
	return sum
}

func dfs1302(root *TreeNode, level int, maxLevel *int, sum *int) {
	if root == nil {
		return
	}
	level++
	if level > *maxLevel {
		*maxLevel = level
		*sum = root.Val
	} else if level == *maxLevel {
		*sum += root.Val
	}
	dfs1302(root.Left, level, maxLevel, sum)
	dfs1302(root.Right, level, maxLevel, sum)
}
