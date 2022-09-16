/**
 * @author: baowj
 * @time: 2022/8/25 16:32:39
 */
package main

import "sort"

func canBeEqual(target []int, arr []int) bool {
	if len(target) != len(arr) {
		return false
	}
	sort.Ints(target)
	sort.Ints(arr)
	for i := 0; i < len(target); i++ {
		if target[i] != arr[i] {
			return false
		}
	}
	return true
}
