/**
 * @author: baowj
 * @time: 2022/4/8 09:34:06
 */
package main

import (
	"container/list"
)

/**
* Definition for a Node.
 */

type Node struct {
	Val      int
	Children []*Node
}

func levelOrder(root *Node) [][]int {
	if root == nil {
		return [][]int{}
	}
	var res [][]int
	queue := list.New()
	queue.Init()
	queue.PushBack(root)
	queue.PushBack(nil)
	var level []int
	for queue.Len() > 1 {
		tmp := queue.Front()
		queue.Remove(tmp)
		if tmp.Value == nil {
			res = append(res, level)
			level = []int{}
			queue.PushBack(nil)
		} else {
			level = append(level, tmp.Value.(*Node).Val)
			for i := 0; i < len(tmp.Value.(*Node).Children); i++ {
				queue.PushBack(tmp.Value.(*Node).Children[i])
			}
		}
	}
	if len(level) > 0 {
		res = append(res, level)
	}
	return res
}
