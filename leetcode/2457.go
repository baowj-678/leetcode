/**
 * @author: baowj
 * @time: 2022/11/2 17:37:08
 */
package main

import "fmt"

func makeIntegerBeautiful(n int64, target int) int64 {
	nums := []int64{}
	sum := 0
	for n > 0 {
		num := n % 10
		n /= 10
		nums = append(nums, num)
		sum += int(num)
	}
	if sum <= target {
		return 0
	}
	nums = append(nums, 0)
	src := make([]int64, len(nums))
	copy(src, nums)
	resNums := make([]int64, len(nums))
	i := 0
	for sum > target && i < len(nums) {
		if nums[i] > 0 {
			resNums[i] = 10 - nums[i]
			nums[i] = 10
			sum = update2457(nums)
		}
		i++
	}
	var res int64 = 0
	var pow int64 = 1
	for len(resNums) > 0 && resNums[len(resNums)-1] == 0 {
		resNums = resNums[:len(resNums)-1]
	}
	for len(resNums) > 0 {
		res += resNums[0] * pow
		resNums = resNums[1:]
		pow *= 10
	}
	return res
}

func update2457(nums []int64) int {
	res := 0
	for i := 0; i < len(nums); i++ {
		if nums[i] > 9 {
			nums[i+1] += nums[i] / 10
			nums[i] %= 10
		}
		res += int(nums[i])
	}
	return res
}

func main() {
	fmt.Println(makeIntegerBeautiful(1, 1))
}
