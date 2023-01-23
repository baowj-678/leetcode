/**
 * @author: baowj
 * @time: 2023/1/23 16:44:26
 * @email: bwj_678@qq.com
 */
package main

import "sort"

func findingUsersActiveMinutes(logs [][]int, k int) []int {
	userActiveTime := map[int][]int{}
	for _, log := range logs {
		if array, ok := userActiveTime[log[0]]; ok {
			userActiveTime[log[0]] = append(array, log[1])
		} else {
			userActiveTime[log[0]] = []int{log[1]}
		}
	}

	res := make([]int, k)
	for _, v := range userActiveTime {
		sort.Ints(v)
		tmp := 0
		last := -1
		for _, time := range v {
			if time != last {
				tmp++
			}
			last = time
		}
		res[tmp-1]++
	}
	return res
}
