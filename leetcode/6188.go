/**
 * @author: baowj
 * @time: 2022/9/25 20:24:21
 */
package main

import "sort"

func sortPeople(names []string, heights []int) []string {
	m := map[int]string{}
	for i := 0; i < len(names); i++ {
		m[heights[i]] = names[i]
	}
	sort.Ints(heights)
	for i := 0; i < len(names); i++ {
		names[i] = m[heights[len(names)-1-i]]
	}
	return names
}
