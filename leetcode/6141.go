/**
 * @author: baowj
 * @time: 2022/8/7 10:47:26
 */
package main

import "sort"

func mergeSimilarItems(items1 [][]int, items2 [][]int) [][]int {
	m := map[int]int{}
	for _, item := range items1 {
		m[item[0]] = item[1]
	}
	for _, item := range items2 {
		m[item[0]] += item[1]
	}
	res := [][]int{}
	for v, w := range m {
		res = append(res, []int{v, w})
	}
	sort.Slice(res, func(i, j int) bool {
		return res[i][0] < res[j][0]
	})
	return res
}
