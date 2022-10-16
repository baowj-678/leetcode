/**
 * @author: baowj
 * @time: 2022/9/25 20:02:36
 */
package main

import (
	"strconv"
)

func rotatedDigits(n int) int {
	res := 0
	for i := 1; i <= n; i++ {
		if canRotated(i) {
			res++
		}
	}
	return res
}

func canRotated(x int) bool {
	s := strconv.Itoa(x)
	isDifferent := false
	for i := 0; i < len(s); i++ {
		switch s[i] {
		case '2':
			isDifferent = true
		case '3':
			return false
		case '4':
			return false
		case '5':
			isDifferent = true
		case '6':
			isDifferent = true
		case '7':
			return false
		case '9':
			isDifferent = true
		}
	}
	return isDifferent
}
