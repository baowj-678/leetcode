/**
 * @author: baowj
 * @time: 2023/9/18 10:22
 */
package main

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func reverseEvenLengthGroups(head *ListNode) *ListNode {
	reverse := func(node *ListNode) (h *ListNode, t *ListNode) {
		t = node
		h = nil
		for node != nil {
			next := node.Next

			node.Next = h
			h = node
			node = next
		}
		return h, t
	}

	node := head
	length := 1
	cnt := 0
	h, t := head, (*ListNode)(nil)
	for node != nil {
		if cnt == 0 {
			h = node
		}
		cnt++
		if cnt == length {
			next := node.Next
			if cnt%2 == 0 {
				node.Next = nil
				sh, st := reverse(h)
				t.Next = sh
				t = st
			} else {
				if t != nil {
					t.Next = h
				}
				t = node
			}
			node = next

			cnt = 0
			length++
		} else {
			node = node.Next
		}
	}
	if cnt == 0 {
		return head
	}
	if cnt%2 == 0 {
		sh, st := reverse(h)
		t.Next = sh
		t = st
	} else {
		t.Next = h
	}
	return head
}
