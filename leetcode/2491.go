/**
 * @author: baowj
 * @time: 2022/12/5 23:16:21
 * @email: bwj_678@qq.com
 */
package main

import "sort"

func dividePlayers(skill []int) int64 {
	sum, res := -1, int64(0)
	sort.Ints(skill)
	sum = skill[0] + skill[len(skill)-1]
	for i := 0; i < len(skill)/2; i++ {
		if sum != skill[i]+skill[len(skill)-1-i] {
			return -1
		}
		res += int64(skill[i] * skill[len(skill)-1-i])
	}
	return res
}
