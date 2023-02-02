/**
 * @author: baowj
 * @time: 2023/2/2 21:01:55
 * @email: bwj_678@qq.com
 */
package main

func countAsterisks(s string) int {
	res := 0
	verticalLineCnt := 0
	for _, c := range s {
		if c == '|' {
			verticalLineCnt++
		} else if c == '*' && verticalLineCnt%2 == 1 {
			res++
		}
	}
	return res
}
