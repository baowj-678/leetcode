/**
 * @author: baowj
 * @time: 2022/7/25 10:32:26
 */
package main

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
type CBTInserter struct {
	root *TreeNode
	heap []*TreeNode
}

func Constructor919(root *TreeNode) CBTInserter {
	cbt := CBTInserter{
		root: root,
		heap: []*TreeNode{nil, root},
	}
	cbt.init(root)
	return cbt
}

func (this *CBTInserter) init(root *TreeNode) {
	if root == nil {
		return
	}
	queue := []*TreeNode{
		root.Left,
		root.Right,
	}
	for len(queue) > 0 {
		node := queue[0]
		queue = queue[1:]
		if node == nil {
			continue
		}
		index := len(this.heap)
		if index%2 == 0 {
			this.heap[index/2].Left = node
		} else {
			this.heap[index/2].Right = node
		}
		this.heap = append(this.heap, node)
		queue = append(queue, node.Left, node.Right)
	}
}

func (this *CBTInserter) Insert(val int) int {
	node := TreeNode{
		Val:   val,
		Left:  nil,
		Right: nil,
	}
	index := len(this.heap)
	if index%2 == 0 {
		this.heap[index/2].Left = &node
	} else {
		this.heap[index/2].Right = &node
	}
	this.heap = append(this.heap, &node)
	return this.heap[index/2].Val
}

func (this *CBTInserter) Get_root() *TreeNode {
	return this.root
}

/**
 * Your CBTInserter object will be instantiated and called as such:
 * obj := Constructor(root);
 * param_1 := obj.Insert(val);
 * param_2 := obj.Get_root();
 */
