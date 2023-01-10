/**
 * @author: baowj
 * @time: 2022/11/16 22:30:44
 * @email: bwj_678@qq.com
 */
package main

func isIdealPermutation(nums []int) bool {
	max, lastMax := -1, -1
	isSame := false
	for i, num := range nums {
		if num > max {
			isSame = false
		} else if num == max {
			isSame = true
		} else {
			if nums[i-1] >= num && nums[i-1] < max {
				return false
			}
			if isSame {
				return false
			}
			if num < lastMax {
				return false
			}
			isSame = false
		}

		lastMax = max
		if num > max {
			max = num
		}
	}
	return true
}
