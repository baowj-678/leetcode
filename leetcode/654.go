/**
 * @author: baowj
 * @time: 2022/8/20 10:19:20
 */
package main

import (
	"math"
)

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func constructMaximumBinaryTree(nums []int) *TreeNode {
	return buildTree(nums, 0, len(nums))
}

func buildTree(nums []int, from int, to int) *TreeNode {
	if from == to {
		return nil
	}
	max, maxIdx := math.MinInt, 0
	for i := from; i < to; i++ {
		if max < nums[i] {
			max = nums[i]
			maxIdx = i
		}
	}
	node := TreeNode{
		Val:   max,
		Left:  nil,
		Right: nil,
	}
	node.Left = buildTree(nums, from, maxIdx)
	node.Right = buildTree(nums, maxIdx+1, to)
	return &node
}
