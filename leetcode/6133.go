/**
 * @author: baowj
 * @time: 2022/7/31 10:37:02
 */
package main

import "math"

func maximumGroups(grades []int) int {
	length := len(grades) * 2
	n := int(math.Floor(math.Sqrt(float64(length))))
	if n*(n+1) <= length {
		return n
	} else if n*(n-1) <= length {
		return n - 1
	} else {
		return n - 2
	}
}

func main() {
	nums := []int{10, 6, 12, 7, 3, 5}
	println(maximumGroups(nums))
}
