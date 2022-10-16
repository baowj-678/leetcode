/**
 * @author: baowj
 * @time: 2022/9/26 10:00:00
 */
package main

import "fmt"

func missingTwo(nums []int) []int {
	nums = append(nums, 0)
	nums = append(nums, 0)
	for i := 0; i < len(nums); i++ {
		j := nums[i] - 1
		for j >= 0 && nums[j] != -1 && nums[j] != j+1 {
			j, nums[j] = nums[j]-1, j+1
		}
	}
	res := []int{}
	for i := 0; i < len(nums); i++ {
		if nums[i] != i+1 {
			res = append(res, i+1)
		}
	}
	return res
}

func main() {
	nums := []int{2, 3}
	fmt.Println(missingTwo(nums))
}
