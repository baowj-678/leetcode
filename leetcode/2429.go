/**
 * @author: baowj
 * @time: 2022/10/3 16:28:05
 */
package main

import (
	"strconv"
)

func minimizeXor(num1 int, num2 int) int {
	cnt := 0
	for num2 > 0 {
		if num2&1 == 1 {
			cnt++
		}
		num2 >>= 1
	}
	a := strconv.FormatInt(int64(num1), 2)
	res := make([]byte, 31)
	start := len(res) - len(a)
	for i := 0; i < len(a); i++ {
		if a[i] == '1' && cnt > 0 {
			res[start+i] = 1
			cnt--
		}
	}
	for i := len(a) - 1; i >= 0; i-- {
		if a[i] == '0' && cnt > 0 {
			res[start+i] = 1
			cnt--
		}
	}
	start--
	for cnt > 0 {
		res[start] = 1
		start--
		cnt--
	}
	ret := 0
	for i := len(res) - 1; i >= 0; i-- {
		if res[i] == 1 {
			ret += 1 << (len(res) - i - 1)
		}
	}
	return ret
}
