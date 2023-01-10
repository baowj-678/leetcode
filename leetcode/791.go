package main

import "sort"

func customSortString(order string, s string) string {
	orderIdx := make([]int, 26)
	for i, x := range order {
		orderIdx[x-'a'] = i
	}
	res := []byte(s)
	sort.Slice(res, func(i, j int) bool {
		return orderIdx[res[i]-'a'] < orderIdx[res[j]-'a']
	})
	return string(res)
}
