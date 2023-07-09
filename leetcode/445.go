/**
 * @author: baowj
 * @time: 2023/7/9 14:54
 */
package main

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	s1, s2 := []int{}, []int{}
	for l1 != nil {
		s1 = append(s1, l1.Val)
		l1 = l1.Next
	}

	for l2 != nil {
		s2 = append(s2, l2.Val)
		l2 = l2.Next
	}

	increament := 0
	var res *ListNode = nil
	for len(s1) > 0 || len(s2) > 0 {
		var v1, v2 int = 0, 0
		if len(s1) > 0 {
			v1 = s1[len(s1)-1]
			s1 = s1[:len(s1)-1]
		}

		if len(s2) > 0 {
			v2 = s2[len(s2)-1]
			s2 = s2[:len(s2)-1]
		}

		increament = increament + v1 + v2
		tmp := &ListNode{
			Val:  increament % 10,
			Next: res,
		}
		increament /= 10
		res = tmp
	}
	if increament > 0 {
		tmp := &ListNode{
			Val:  increament,
			Next: res,
		}
		res = tmp
	}
	return res
}
