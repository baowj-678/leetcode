/**
 * @author: baowj
 * @time: 2022/12/15 22:20:02
 * @email: bwj_678@qq.com
 */
package main

import "math"

func beautySum(s string) int {
	prefix := make([][26]int, len(s))
	for i, c := range s {
		if i > 0 {
			for k := 0; k < 26; k++ {
				prefix[i][k] = prefix[i-1][k]
			}
		}
		prefix[i][c-'a']++
	}
	res := 0
	for l := 0; l < len(s); l++ {
		for r := l + 1; r < len(s); r++ {
			max, min := 0, math.MaxInt
			lP, rP := [26]int{}, prefix[r]
			if l > 0 {
				lP = prefix[l-1]
			}
			for k := 0; k < 26; k++ {
				if rP[k]-lP[k] > max {
					max = rP[k] - lP[k]
				}
				if rP[k]-lP[k] < min && rP[k]-lP[k] > 0 {
					min = rP[k] - lP[k]
				}
			}
			res += max - min
		}
	}
	return res
}
