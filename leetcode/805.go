/**
 * @author: baowj
 * @time: 2022/11/16 23:32:34
 * @email: bwj_678@qq.com
 */
package main

import "sort"

func splitArraySameAverage(nums []int) bool {
	sum := 0
	for _, num := range nums {
		sum += num
	}
	end := len(nums) - 1
	for i, num := range nums {
		if num*len(nums) == sum {
			nums[end], nums[i] = 0, nums[end]
			end--
		}
	}
	if end < 0 {
		return true
	}
	nums = nums[:end+1]
	sort.Ints(nums)
	return dfs805(nums, 0, 0, 0, sum)
}

func dfs805(nums []int, idx, tmp, cnt, sum int) bool {
	if tmp == sum {
		return false
	}
	if tmp*len(nums) == sum*cnt && cnt != 0 {
		return true
	}
	if tmp*len(nums) > sum*cnt && nums[idx]*len(nums) >= sum {
		return false
	}

	if idx < len(nums)-1 && dfs805(nums, idx+1, tmp, cnt, sum) {
		return true
	}

	tmp += nums[idx]
	cnt++
	if tmp == sum {
		return false
	}
	if tmp*len(nums) == sum*cnt && cnt != 0 {
		return true
	}
	if tmp*len(nums) > sum*cnt && nums[idx]*len(nums) >= sum {
		return false
	}
	if idx < len(nums)-1 && dfs805(nums, idx+1, tmp, cnt, sum) {
		return true
	}
	return false
}

func main() {
	nums := []int{3, 12, 3, 4, 5, 6}
	print(splitArraySameAverage(nums))
}
