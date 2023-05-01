/**
 * @author: baowj
 * @time: 2023/4/30 09:51:14
 * @email: bwj_678@qq.com
 */
package main

func numMovesStones(a int, b int, c int) []int {
	if a > b {
		a, b = b, a
	}
	if a < c && c < b {
		c, b = b, c
	} else if c < a {
		a, b, c = c, a, b
	}
	res := make([]int, 2)
	if b-a > 1 {
		if c-b > 1 {
			if b-a == 2 || c-b == 2 {
				res[0] = 1
			} else {
				res[0] = 2
			}
		} else {
			res[0] = 1
		}
	} else {
		if c-b > 1 {
			res[0] = 1
		} else {
			res[0] = 0
		}
	}

	res[1] = c - b + b - a - 2
	return res
}
