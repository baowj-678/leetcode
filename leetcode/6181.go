/**
 * @author: baowj
 * @time: 2022/9/18 13:32:58
 */
package main

func longestContinuousSubstring(s string) int {
	res := 0
	var prevC int32 = 0
	tmpLength := 1
	for _, c := range s {
		if c-prevC == 1 {
			tmpLength++
		} else {
			tmpLength = 1
		}
		prevC = c
		if tmpLength > res {
			res = tmpLength
		}
	}
	return res
}
