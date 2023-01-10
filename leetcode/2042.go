/**
 * @author: baowj
 * @time: 2023/1/8 12:18:25
 * @email: bwj_678@qq.com
 */
package main

import (
	"strconv"
	"strings"
)

func areNumbersAscending(s string) bool {
	tokens := strings.Split(s, " ")
	lastNum := 0
	for _, token := range tokens {
		if token[0] >= '0' && token[0] <= '9' {
			n, _ := strconv.Atoi(token)
			if n <= lastNum {
				return false
			}
			lastNum = n
		}
	}
	return true
}
