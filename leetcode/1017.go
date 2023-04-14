/**
 * @author: baowj
 * @time: 2023/4/6 16:21
 */
package main

import "strconv"

func baseNeg2(n int) string {
	if n == 0 {
		return "0"
	}
	nums := make([]int, 64)
	i := 0
	for n != 0 {
		if -n%2 != 0 {
			nums[i] = 1
		}
		n -= nums[i]
		n /= -2
		i++
	}

	res := ""
	for j := i - 1; j >= 0; j-- {
		res += strconv.Itoa(nums[j])
	}
	return res
}
