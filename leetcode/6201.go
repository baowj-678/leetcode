/**
 * @author: baowj
 * @time: 2022/10/9 15:51:43
 */
package main

func findArray(pref []int) []int {
	res := make([]int, len(pref))
	res[0] = pref[0]
	for i := 1; i < len(pref); i++ {
		res[i] = pref[i-1] ^ pref[i]
	}
	return res
}
