/**
 * @author: baowj
 * @time: 2022/11/19 21:46:25
 * @email: bwj_678@qq.com
 */
package main

func largestAltitude(gain []int) int {
	max, tmp := 0, 0
	for _, inc := range gain {
		tmp += inc
		if tmp > max {
			max = tmp
		}
	}
	return max
}
