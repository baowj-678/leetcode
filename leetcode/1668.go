/**
 * @author: baowj
 * @time: 2022/11/6 18:56:15
 */
package main

import "strings"

func maxRepeating(sequence string, word string) int {
	str := ""
	res := 0
	for len(str) <= len(sequence) {
		str += word
		res++
		if !strings.Contains(sequence, str) {
			return res - 1
		}
	}
	return res
}
