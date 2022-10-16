/**
 * @author: baowj
 * @time: 2022/10/16 19:11:36
 */
package main

import "strconv"

func sumOfNumberAndReverse(num int) bool {
	for i := 0; i <= num; i++ {
		rev := i
		str := strconv.Itoa(rev)
		strbyte := []byte(str)
		for l, r := 0, len(strbyte)-1; l < r; l, r = l+1, r-1 {
			strbyte[l], strbyte[r] = strbyte[r], strbyte[l]
		}
		rev, _ = strconv.Atoi(string(strbyte))
		if i+rev == num {
			return true
		}
	}
	return false
}
