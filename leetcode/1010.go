/**
 * @author: baowj
 * @time: 2023/5/8 9:51
 */
package main

func numPairsDivisibleBy60(time []int) int {
	times := make([]int, 60)
	res := 0
	for _, t := range time {
		times[t%60]++
	}

	for i := 1; i <= 29; i++ {
		res += times[i] * times[60-i]
	}
	res += times[0] * (times[0] - 1) / 2
	res += times[30] * (times[30] - 1) / 2
	return res
}
