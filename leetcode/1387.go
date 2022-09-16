/**
 * @author: baowj
 * @time: 2022/9/9 21:56:35
 */
package main

import "sort"

func getKth(lo int, hi int, k int) int {
	m := map[int]int{}
	array := make([][]int, hi-lo+1)
	for i := lo; i <= hi; i++ {
		array[i-lo] = []int{i, getWeight(i, m)}
	}
	sort.Slice(array, func(i, j int) bool {
		if array[i][1] < array[j][1] {
			return true
		} else if array[i][1] > array[j][1] {
			return false
		} else {
			return array[i][0] < array[j][0]
		}
	})
	return array[k-1][0]
}

func getWeight(x int, m map[int]int) int {
	cnt := 0
	if x == 1 {
		return 0
	}
	if v, ok := m[x]; ok {
		cnt += v
	} else {
		if x%2 == 0 {
			cnt += getWeight(x/2, m)
			m[x] = cnt
		} else {
			cnt += getWeight(3*x+1, m)
			m[x] = cnt
		}
	}
	return cnt + 1
}
