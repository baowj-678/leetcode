/**
 * @author: baowj
 * @time: 2023/2/14 11:17
 */
package main

import "fmt"

func longestWPI(hours []int) int {
	sum := make([]int, len(hours)+1)
	sum[0] = 0
	for i, hour := range hours {
		if hour > 8 {
			sum[i+1] = sum[i] + 1
		} else {
			sum[i+1] = sum[i] - 1
		}
	}
	stk := []int{0}
	for i, num := range sum {
		if num < sum[stk[len(stk)-1]] {
			stk = append(stk, i)
		}
	}

	fmt.Println(sum)
	fmt.Println(stk)
	res := 0
	for i := len(sum) - 1; i >= 1; i-- {
		idx := i
		for len(stk) > 0 && sum[i] > sum[stk[len(stk)-1]] {
			idx = stk[len(stk)-1]
			stk = stk[:len(stk)-1]
		}
		if res < i-idx {
			res = i - idx
		}
	}
	return res
}

func main() {
	hours := []int{6, 9, 9}
	print(longestWPI(hours))
}
