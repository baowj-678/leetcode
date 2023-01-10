/**
 * @author: baowj
 * @time: 2022/12/27 21:16:40
 * @email: bwj_678@qq.com
 */
package main

func minimumMoves(s string) int {
	res := 0
	for i := 0; i < len(s); i++ {
		c := s[i]
		if c == 'X' {
			res++
			i += 2
		}
	}
	return res
}
