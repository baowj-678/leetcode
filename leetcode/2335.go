/**
 * @author: baowj
 * @time: 2023/2/11 12:32:49
 * @email: bwj_678@qq.com
 */
package main

import "sort"

func fillCups(amount []int) int {
	sort.Ints(amount)
	if amount[2] >= amount[0]+amount[1] {
		return amount[2]
	} else {
		res := amount[2]
		amount[2] -= amount[1] - amount[0]
		amount[1] = amount[0]
		amount[0] -= amount[2] / 2
		amount[1] -= (amount[2] - amount[2]/2)
		if amount[0] > amount[1] {
			res += amount[0]
		} else {
			res += amount[1]
		}
		return res
	}
}
