/**
 * @author: baowj
 * @time: 2022/9/11 16:02:34
 */
package main

func minGroups(intervals [][]int) int {
	array := make([]int, 1000000+50)
	for _, interval := range intervals {
		array[interval[0]]++
		array[interval[1]+1]--
	}
	max, tmp := 0, 0
	for _, x := range array {
		tmp += x
		if max < tmp {
			max = tmp
		}
	}
	return max
}
