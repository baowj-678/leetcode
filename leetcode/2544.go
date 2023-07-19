/**
 * @author: baowj
 * @time: 2023/7/12 19:18
 */
package main

func alternateDigitSum(n int) int {
	nums := []int{}
	for n > 0 {
		nums = append(nums, n%10)
		n /= 10
	}
	flag := true
	res := 0
	for i := len(nums) - 1; i >= 0; i-- {
		num := nums[i]
		if flag {
			res += num
		} else {
			res -= num
		}
		flag = !flag
	}
	return res
}
