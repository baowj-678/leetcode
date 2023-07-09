/**
 * @author: baowj
 * @time: 2023/7/9 15:16
 */
package main

import "sort"

func maximumEvenSplit(finalSum int64) []int64 {
	if finalSum%2 > 0 {
		return []int64{}
	}
	finalSum /= 2

	v := sort.Search(1e6, func(i int) bool {
		return ((1+int64(i))*int64(i))/2 > finalSum
	})

	finalSum *= 2
	var sum int64 = 0
	res := make([]int64, v-1)
	for i := 0; i < v-1; i++ {
		sum += (int64(i) + 1) * 2
		res[i] = (int64(i) + 1) * 2
	}
	if sum < finalSum {
		res[len(res)-1] += (finalSum - sum)
	}
	return res
}
