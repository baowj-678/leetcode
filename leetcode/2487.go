package main

func removeNodes(head *ListNode) *ListNode {
	stack := []*ListNode{}
	for head != nil {
		for len(stack) > 0 {
			top := stack[len(stack)-1]
			if top.Val < head.Val {
				stack = stack[:len(stack)-1]
			} else {
				break
			}
		}
		stack = append(stack, head)
		head = head.Next
	}
	for i, node := range stack {
		if i+1 < len(stack) {
			node.Next = stack[i+1]
		}
	}
	stack[len(stack)-1].Next = nil
	head = stack[0]
	return head
}
