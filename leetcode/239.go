/**
 * @author: baowj
 * @time: 2022/8/17 10:14:17
 */
package main

import (
	"fmt"
)

func maxSlidingWindow(nums []int, k int) []int {
	queue := []int{}
	res := make([]int, len(nums)-k+1)
	push := func(idx int) {
		for len(queue) > 0 && nums[queue[len(queue)-1]] <= nums[idx] {
			queue = queue[:len(queue)-1]
		}
		queue = append(queue, idx)
	}
	for i, _ := range nums {
		push(i)
		for queue[0] < i-k+1 {
			queue = queue[1:]
		}
		if i >= k-1 {
			res[i-k+1] = nums[queue[0]]
		}
	}
	return res
}

func main() {
	nums := []int{1}
	fmt.Println(maxSlidingWindow(nums, 1))
}
