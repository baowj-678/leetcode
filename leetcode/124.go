/**
 * @author: baowj
 * @time: 2022/7/25 11:00:17
 */
package main

const MIN = -100000000

func maxPathSum(root *TreeNode) int {
	max, _ := dfs(root)
	return max
}

func dfs(root *TreeNode) (max int, nodeMax int) {
	max, nodeMax, leftNodeMax, rightMax, rightNodeMax := MIN, MIN, MIN, MIN, MIN
	if root.Left != nil {
		max, leftNodeMax = dfs(root.Left)
	}
	if root.Right != nil {
		rightMax, rightNodeMax = dfs(root.Right)
		if rightMax > max {
			max = rightMax
		}
	}

	if nodeMax < leftNodeMax+root.Val {
		nodeMax = leftNodeMax + root.Val
	}
	if nodeMax < rightNodeMax+root.Val {
		nodeMax = rightNodeMax + root.Val
	}
	if max < leftNodeMax+rightNodeMax+root.Val {
		max = leftNodeMax + rightNodeMax + root.Val
	}
	if nodeMax < root.Val {
		nodeMax = root.Val
	}

	if max < nodeMax {
		max = nodeMax
	}
	return max, nodeMax
}
