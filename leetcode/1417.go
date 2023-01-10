/**
 * @author: baowj
 * @time: 2022/8/11 10:24:02
 */
package main

import (
	"unicode"
)

func reformat(s string) string {
	alphaP, numP := 0, 0
	for i := 0; i < len(s); i++ {
		if unicode.IsDigit(rune(s[i])) {
			numP += 1
		} else {
			alphaP += 1
		}
	}
	if abs1417(numP, alphaP) > 1 {
		return ""
	}
	res := make([]int32, len(s))
	if alphaP >= numP {
		alphaP = 0
		for _, c := range s {
			if unicode.IsLetter(c) {
				res[alphaP] = c
				alphaP += 2
			}
		}
		numP = 1
		for _, c := range s {
			if unicode.IsDigit(c) {
				res[numP] = c
				numP += 2
			}
		}
	} else {
		alphaP = 1
		for _, c := range s {
			if unicode.IsLetter(c) {
				res[alphaP] = c
				alphaP += 2
			}
		}
		numP = 0
		for _, c := range s {
			if unicode.IsDigit(c) {
				res[numP] = c
				numP += 2
			}
		}
	}
	return string(res)
}

func abs1417(x int, y int) int {
	if x > y {
		return x - y
	} else {
		return y - x
	}
}
