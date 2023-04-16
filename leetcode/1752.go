/**
 * @author: baowj
 * @time: 2022/11/27 10:29:40
 * @email: bwj_678@qq.com
 */
package main

func check1752(nums []int) bool {
	for i := 0; i < len(nums); i++ {
		isOk := true
		last := nums[i]
		for j := 1; j < len(nums); j++ {
			if nums[(i+j)%len(nums)] < last {
				isOk = false
				break
			}
			last = nums[(i+j)%len(nums)]
		}
		if isOk {
			return true
		}
	}
	return false
}
