/**
 * @author: baowj
 * @time: 17:25
 */

package main

func numComponents(head *ListNode, nums []int) int {
	m := map[int]bool{}
	for _, num := range nums {
		m[num] = true
	}
	res := 0
	lastIncluded := false
	for head != nil {
		if _, ok := m[head.Val]; ok {
			if !lastIncluded {
				res++
				lastIncluded = true
			}
		} else {
			lastIncluded = false
		}
		head = head.Next
	}
	return res
}
