/**
 * @author: baowj
 * @time: 2023/2/12 16:40:00
 * @email: bwj_678@qq.com
 */
package main

import "sort"

func countFairPairs(nums []int, lower int, upper int) int64 {
	lowerNums, upperNums := make([]int, len(nums)), make([]int, len(nums))
	sort.Ints(nums)
	for i := 0; i < len(nums); i++ {
		lowerNums[i] = lower - nums[i]
		upperNums[i] = upper - nums[i]
	}
	var res int64 = 0
	for _, num := range nums {
		l := sort.Search(len(nums), func(i int) bool {
			return lowerNums[i] <= num
		})

		r := sort.Search(len(upperNums), func(i int) bool {
			return upperNums[i] < num
		})

		length := r - l
		if length <= 0 {
			continue
		}
		if num*2 >= lower && num*2 <= upper {
			length--
		}
		res += int64(length)
	}
	return res / 2
}
