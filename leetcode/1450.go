/**
 * @author: baowj
 * @time: 2022/8/19 20:57:16
 */
package main

func busyStudent(startTime []int, endTime []int, queryTime int) int {
	queryArray := [1000 + 5]int{0}
	for i := 0; i < len(startTime); i++ {
		queryArray[startTime[i]]++
		queryArray[endTime[i]+1]--
	}
	for i := 1; i <= queryTime; i++ {
		queryArray[i] += queryArray[i-1]
	}
	return queryArray[queryTime]
}
