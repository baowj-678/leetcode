/**
 * @author: baowj
 * @time: 2023/9/20 9:56
 */
package main

func minCount(coins []int) int {
	res := 0
	for _, coin := range coins {
		res += (coin + 1) / 2
	}
	return res
}
