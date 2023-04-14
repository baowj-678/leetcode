/**
 * @author: baowj
 * @time: 2023/4/10 10:05
 */
package main

func nextLargerNodes(head *ListNode) []int {
	queue := [][2]int{}
	res := []int{}
	this := head
	for this != nil {
		v := this.Val
		for i := len(queue) - 1; i >= 0; i-- {
			if queue[i][0] < v {
				h := queue[len(queue)-1]
				queue = queue[:len(queue)-1]

				res[h[1]] = v
			}
		}

		queue = append(queue, [2]int{v, len(res)})
		res = append(res, 0)
		this = this.Next
	}
	for _, q := range queue {
		res[q[1]] = 0
	}
	return res
}
