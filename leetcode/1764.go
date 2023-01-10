/**
 * @author: baowj
 * @time: 2022/12/17 19:41:07
 * @email: bwj_678@qq.com
 */
package main

func canChoose(groups [][]int, nums []int) bool {
	idx := 0
	for _, group := range groups {
		can := false
		for i := idx; i < len(nums); i++ {
			isOK := true
			for k := 0; k < len(group); k++ {
				if k+i >= len(nums) {
					return false
				}
				if group[k] != nums[i+k] {
					isOK = false
					break
				}
			}
			if isOK {
				can = true
				idx = i + len(group)
				break
			}
		}
		if !can {
			return false
		}
	}
	return true
}
