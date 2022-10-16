/**
 * @author: baowj
 * @time: 2022/10/16 19:02:56
 */
package main

import (
	"strconv"
)

func countDistinctIntegers(nums []int) int {
	m := map[int]bool{}
	for _, num := range nums {
		m[num] = true
		str := strconv.Itoa(num)
		strbyte := []byte(str)
		for l, r := 0, len(strbyte)-1; l < r; l, r = l+1, r-1 {
			strbyte[l], strbyte[r] = strbyte[r], strbyte[l]
		}
		num, _ = strconv.Atoi(string(strbyte))
		m[num] = true
	}
	return len(m)
}
