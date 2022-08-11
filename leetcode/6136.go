/**
 * @author: baowj
 * @time: 2022/8/7 17:28:00
 */
package main

func arithmeticTriplets(nums []int, diff int) int {
	cnt := 0
	for i := 0; i < len(nums); i++ {
		for j := i + 1; j < len(nums); j++ {
			if nums[j]-nums[i] > diff {
				break
			} else if nums[j]-nums[i] < diff {
				continue
			} else {
				for k := j + 1; k < len(nums); k++ {
					if nums[k]-nums[j] > diff {
						break
					} else if nums[k]-nums[j] < diff {
						continue
					} else {
						cnt += 1
						break
					}
				}
			}
		}
	}
	return cnt
}
