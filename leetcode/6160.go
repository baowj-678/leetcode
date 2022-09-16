/**
 * @author: baowj
 * @time: 2022/8/28 21:21:07
 */
package main

import (
	"sort"
)

func answerQueries(nums []int, queries []int) []int {
	sort.Ints(nums)
	for i := 1; i < len(nums); i++ {
		nums[i] += nums[i-1]
	}
	res := make([]int, len(queries))
	for i := 0; i < len(res); i++ {
		idx := sort.Search(len(nums), func(j int) bool { return nums[j] > queries[i] })
		res[i] = idx
	}
	return res
}
