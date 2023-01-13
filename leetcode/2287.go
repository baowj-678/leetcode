/**
 * @author: baowj
 * @time: 2023/1/13 11:37:59
 * @email: bwj_678@qq.com
 */
package main

import "math"

func rearrangeCharacters(s string, target string) int {
	targetCnt := [26]int{}
	for _, c := range target {
		targetCnt[c-'a']++
	}
	sCnt := [26]int{}
	for _, c := range s {
		sCnt[c-'a']++
	}
	res := math.MaxInt
	for i := 0; i < 26; i++ {
		if targetCnt[i] != 0 {
			x := sCnt[i] / targetCnt[i]
			if x < res {
				res = x
			}
		}
	}
	return res
}
