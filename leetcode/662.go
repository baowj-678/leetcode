/**
 * @author: baowj
 * @time: 2022/8/27 11:10:51
 */
package main

import (
	"fmt"
	"math"
)

func widthOfBinaryTree(root *TreeNode) int {
	return dfs662(root, 0, 0, &[]int{})
}

func dfs662(root *TreeNode, idx int, level int, levelLeftIdx *[]int) int {
	fmt.Println()
	if root == nil {
		return 0
	}
	if len(*levelLeftIdx) <= level {
		*levelLeftIdx = append(*levelLeftIdx, math.MaxInt)
	}
	if (*levelLeftIdx)[level] > idx {
		(*levelLeftIdx)[level] = idx
	}
	maxWidth := idx - (*levelLeftIdx)[level]
	width := dfs662(root.Left, idx*2, level+1, levelLeftIdx)
	if maxWidth < width {
		maxWidth = width
	}
	width = dfs662(root.Right, idx*2+1, level+1, levelLeftIdx)
	if maxWidth < width {
		maxWidth = width
	}
	return maxWidth
}
