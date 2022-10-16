/**
 * @author: baowj
 * @time: 2022/9/19 09:48:07
 */
package main

import "sort"

func frequencySort(nums []int) []int {
	m := map[int]int{}
	for _, x := range nums {
		if v, ok := m[x]; ok {
			m[x] = v + 1
		} else {
			m[x] = 1
		}
	}
	array := [][]int{}
	for k, v := range m {
		array = append(array, []int{k, v})
	}
	res := make([]int, len(nums))
	sort.Slice(array, func(i, j int) bool {
		if array[i][1] < array[j][1] {
			return true
		} else if array[i][1] > array[j][1] {
			return false
		} else {
			return array[i][0] > array[j][0]
		}
	})
	k := 0
	for i := 0; i < len(array); i++ {
		res[k] = array[i][0]
		array[i][1]--
		k++
		if array[i][1] > 0 {
			i--
		}
	}
	return res
}
