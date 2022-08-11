/**
 * @author: baowj
 * @time: 2022/8/7 17:26:20
 */
package main

func validPartition(nums []int) bool {
	flag := make([]bool, 3)
	if nums[0] == nums[1] {
		flag[2] = true
	}
	flag[1] = false
	flag[0] = true
	for i := 2; i < len(nums); i++ {
		tmp := false
		if nums[i] == nums[i-1] && nums[i-1] == nums[i-2] && flag[0] {
			tmp = true
		} else if nums[i] == nums[i-1] && flag[1] {
			tmp = true
		} else if nums[i]-nums[i-1] == 1 && nums[i-1]-nums[i-2] == 1 && flag[0] {
			tmp = true
		}
		flag[0] = flag[1]
		flag[1] = flag[2]
		flag[2] = tmp
	}
	return flag[2]
}
