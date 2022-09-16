/**
 * @author: baowj
 * @time: 2022/9/14 22:02:28
 */
package main

import "sort"

func trimMean(arr []int) float64 {
	sort.Ints(arr)
	start := len(arr) / 20
	end := len(arr) / 20 * 19
	sum := 0
	for i := start; i < end; i++ {
		sum += arr[i]
	}
	return float64(sum) / float64(end-start)
}
