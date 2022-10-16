/**
 * @author: baowj
 * @time: 2022/9/17 20:52:11
 */
package main

func longestNiceSubarray(nums []int) int {
	res := 1
	for i := 0; i < len(nums); i++ {
		for j := i + 1; j < len(nums); j++ {
			isNice := true
			for k := i; k < j; k++ {
				if nums[k]&nums[j] != 0 {
					isNice = false
					break
				}
			}
			if isNice {
				if j-i+1 > res {
					res = j - i + 1
				}
			} else {
				break
			}
		}
	}
	return res
}
