/**
 * @author: baowj
 * @time: 2022/10/16 19:15:39
 */
package main

func countSubarrays(nums []int, minK int, maxK int) int64 {
	minC, maxC := 0, 0
	stop := -1
	l := -1
	var res int64 = 0
	for i, num := range nums {
		if num < minK || num > maxK {
			stop = i
			minC, maxC = 0, 0
			l = -1
		} else {
			if l == -1 {
				l = i
			}
			if num == minK {
				minC++
			}
			if num == maxK {
				maxC++
			}
			for l < i {
				if nums[l] == minK {
					if minC > 1 {
						minC--
					} else {
						break
					}
				}
				if nums[l] == maxK {
					if maxC > 1 {
						maxC--
					} else {
						break
					}
				}
				l++
			}
			if minC > 0 && maxC > 0 {
				res += int64(l - stop)
			}
		}
	}
	return res
}
