package main

func alternatingSubarray(nums []int) int {
	start, flag, res := 0, true, -1
	for i := 1; i < len(nums); i++ {
		num := nums[i]
		if flag && num-nums[i-1] == 1 {
			flag = !flag
			if res < i-start+1 {
				res = i - start + 1
			}
			continue
		}
		if !flag && nums[i-1]-num == 1 {
			flag = !flag
			if res < i-start+1 {
				res = i - start + 1
			}
			continue
		}
		if num-nums[i-1] == 1 {
			start = i - 1
			flag = false
		} else {
			start = i
			flag = true
		}
	}
	return res
}
