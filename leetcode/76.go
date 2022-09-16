/**
 * @author: baowj
 * @time: 2022/8/16 11:43:28
 */
package main

import (
	"math"
	"unicode"
)

func minWindow(s string, t string) string {
	tCnt := [52]int{}
	sCnt := [52]int{}
	for i := 0; i < 52; i++ {
		tCnt[i] = 0
		sCnt[i] = 0
	}
	for _, c := range t {
		if unicode.IsLower(c) {
			tCnt[c-'a']++
		} else {
			tCnt[c-'A'+26]++
		}
	}
	start, length, end, last := 0, math.MaxInt, 0, -1
	for i, c := range s {
		if unicode.IsLower(c) {
			sCnt[c-'a']++
		} else {
			sCnt[c-'A'+26]++
		}
		for cmp76(sCnt, tCnt) {
			if i-last+1 < length {
				start = last
				end = i
				length = end - start + 1
			}
			lastC := s[last]
			last++
			if unicode.IsLower(rune(lastC)) {
				sCnt[lastC-'a']--
			} else {
				sCnt[lastC-'A'+26]--
			}
		}
	}
	return s[start : end+1]
}

func cmp76(x [52]int, y [52]int) bool {
	for i := 0; i < 52; i++ {
		if x[i] < y[i] {
			return false
		}
	}
	return true
}
