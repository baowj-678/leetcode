/**
 * @author: baowj
 * @time: 2022/4/19 16:27:55
 */
package main

import (
	"math"
)

func shortestToChar(s string, c byte) []int {
	ans := make([]int, len(s))
	lastIdx := -1
	for i := 0; i < len(s); i++ {
		if s[i] == c {
			ans[i] = 0
			lastIdx = i
		} else {
			if lastIdx != -1 {
				ans[i] = i - lastIdx
			} else {
				ans[i] = math.MaxInt
			}
		}
	}
	lastIdx = -1
	for i := len(s) - 1; i >= 0; i-- {
		if s[i] == c {
			ans[i] = 0
			lastIdx = i
		} else {
			if lastIdx != -1 {
				if ans[i] > lastIdx-i {
					ans[i] = lastIdx - i
				}
			}
		}
	}
	return ans
}
