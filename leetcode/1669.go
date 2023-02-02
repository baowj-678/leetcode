/**
 * @author: baowj
 * @time: 2023/2/2 21:07:19
 * @email: bwj_678@qq.com
 */
package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func mergeInBetween(list1 *ListNode, a int, b int, list2 *ListNode) *ListNode {
	tail2 := list2
	for tail2 != nil {
		if tail2.Next != nil {
			tail2 = tail2.Next
		} else {
			break
		}
	}
	res := list1
	cnt := 0
	var head1, tail1 *ListNode = nil, nil
	for list1 != nil {
		cnt++
		if cnt == a {
			head1 = list1
		}
		list1 = list1.Next
		if cnt == b {
			tail1 = list1
			break
		}
	}
	head1.Next = list2
	tail2.Next = tail1.Next
	return res
}
