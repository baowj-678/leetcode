/**
 * @author: baowj
 * @time: 2023/6/12 19:10
 */
package main

func removeZeroSumSublists(head *ListNode) *ListNode {
	m := map[int]*ListNode{}
	p := head
	sum := 0
	for p != nil {
		sum += p.Val
		m[sum] = p
		p = p.Next
	}

	if v, ok := m[0]; ok {
		head = v.Next
	}
	p = head
	sum = 0
	for p != nil {
		sum += p.Val
		if v, ok := m[sum]; ok {
			p.Next = v.Next
		}
		p = p.Next
	}
	return head
}
