/**
 * @author: baowj
 * @time: 2023/6/12 18:44
 */
package main

import "sort"

func miceAndCheese(reward1 []int, reward2 []int, k int) int {
	res := 0
	for i := 0; i < len(reward1); i++ {
		reward1[i] -= reward2[i]
		res += reward2[i]
	}
	sort.Ints(reward1)
	for i := 0; i < k; i++ {
		res += reward1[len(reward1)-i-1]
	}
	return res
}
