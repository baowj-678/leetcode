/**
 * @author: baowj
 * @time: 2023/2/4 21:09:48
 * @email: bwj_678@qq.com
 */
package main

import "sort"

func getMaximumConsecutive(coins []int) int {
	max := 0
	sort.Ints(coins)
	for _, coin := range coins {
		if coin <= max+1 {
			max = coin + max
		} else {
			break
		}
	}
	return max + 1
}
