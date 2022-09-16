/**
 * @author: baowj
 * @time: 2022/9/5 11:52:22
 */
package main

import (
	"fmt"
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
func maxSumBST(root *TreeNode) int {
	_, _, maxSum, _, _ := dfs1373(root)
	return maxSum
}

func dfs1373(root *TreeNode) (isBST bool, sum int, maxSum int, max int, min int) {
	if root == nil {
		isBST = true
		sum = 0
		maxSum = 0
		max = math.MinInt
		min = math.MaxInt
		return
	}
	lBST, lsum, lmsum, lmax, lmin := dfs1373(root.Left)
	rBST, rsum, rmsum, rmax, rmin := dfs1373(root.Right)
	isBST = false
	sum = 0
	maxSum = 0
	if lmsum > maxSum {
		maxSum = lmsum
	}
	if rmsum > maxSum {
		maxSum = rmsum
	}
	fmt.Println(root.Val, rsum, lsum)
	if lBST && rBST {
		if root.Left == nil && root.Right == nil {
			isBST = true
			sum = root.Val
		} else if root.Left == nil {
			if root.Val < rmin {
				isBST = true
				sum = root.Val + rsum
			}
		} else if root.Right == nil {
			if lmax < root.Val {
				isBST = true
				sum = root.Val + lsum
			}
		} else {
			if lmax < root.Val && rmin > root.Val {
				isBST = true
				sum = root.Val + lsum + rsum
			}
		}
	}
	min = root.Val
	max = root.Val
	if min > rmin {
		min = rmin
	}
	if min > lmin {
		min = lmin
	}
	if max < rmax {
		max = rmax
	}
	if max < lmax {
		max = lmax
	}
	if sum > maxSum {
		maxSum = sum
	}
	return
}
