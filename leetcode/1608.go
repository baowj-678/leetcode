/**
 * @author: baowj
 * @time: 2022/9/12 10:54:02
 */
package main

func specialArray(nums []int) int {
	for i := 1; i <= 100; i++ {
		cnt := 0
		for j := 0; j < len(nums); j++ {
			if nums[j] >= i {
				cnt++
			}
		}
		if cnt == i {
			return i
		}
	}
	return -1
}
