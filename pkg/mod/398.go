/**
 * @author: baowj
 * @time: 2022/4/25 10:07:37
 */
package main

import "math/rand"

type Solution struct {
	array []int
}

func Constructor1(nums []int) Solution {
	obj := Solution{
		nums,
	}
	return obj
}

func (this *Solution) Pick(target int) int {
	cnt := 0
	ans := 0
	for i, num := range this.array {
		if num == target {
			cnt++
			if rand.Intn(cnt) == 0 {
				ans = i
			}
		}
	}
	return ans
}
