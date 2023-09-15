/**
 * @author: baowj
 * @time: 2022/8/25 15:58:39
 */
package main

import "fmt"

func findClosestElements(arr []int, k int, x int) []int {
	AbsInt := func(x int) int {
		if x < 0 {
			return -x
		}
		return x
	}
	start, end := 0, k-1
	for i := k; i < len(arr); i++ {
		if AbsInt(arr[i]-x) < AbsInt(arr[end]-x) {
			end = i
			start = end - k + 1
		} else if AbsInt(arr[i]-x) >= AbsInt(arr[end]-x) {
			if AbsInt(arr[i]-x) < AbsInt(arr[start]-x) {
				end = i
				start = end - k + 1
			}
		}
	}
	return arr[start : end+1]
}

func main() {
	arr := []int{1, 2, 3, 4, 5}
	fmt.Println(findClosestElements(arr, 4, -1))
}
