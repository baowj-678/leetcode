/**
 * @author: baowj
 * @time: 2022/7/31 09:49:46
 */
package main

import (
	"fmt"
	"math"
)

func maxLevelSum(root *TreeNode) int {
	queue := []*TreeNode{}
	queue = append(queue, root)
	queue = append(queue, nil)
	level, maxSum, maxLevel, tmpSum := 1, math.MinInt, 1, 0
	for len(queue) > 1 {
		element := queue[0]
		queue = queue[1:]
		fmt.Print(element)
		if element == nil {
			if tmpSum > maxSum {
				maxLevel = level
				maxSum = tmpSum
			}
			queue = append(queue, nil)
			level += 1
			tmpSum = 0
		} else {
			tmpSum += element.Val
			if element.Left != nil {
				queue = append(queue, element.Left)
			}
			if element.Right != nil {
				queue = append(queue, element.Right)
			}
		}
	}
	if tmpSum > maxSum {
		maxLevel = level
	}
	return maxLevel
}
