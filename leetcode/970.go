/**
 * @author: baowj
 * @time: 2023/5/2 10:15:30
 * @email: bwj_678@qq.com
 */
package main

import "sort"

func powerfulIntegers(x int, y int, bound int) []int {
	set := map[int]struct{}{}
	xExp, yExp := 1, 1
	for i := 0; i < 21; i++ {
		if xExp >= bound {
			break
		}
		yExp = 1
		for j := 0; j < 21; j++ {
			if xExp+yExp <= bound {
				set[xExp+yExp] = struct{}{}
			} else {
				break
			}
			if y == 1 {
				break
			}
			yExp *= y
		}

		if x == 1 {
			break
		}
		xExp *= x
	}
	res := make([]int, len(set))
	i := 0
	for k, _ := range set {
		res[i] = k
		i++
	}
	sort.Ints(res)
	return res
}
