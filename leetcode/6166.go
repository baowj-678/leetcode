/**
 * @author: baowj
 * @time: 2022/8/21 17:25:38
 */
package main

import "strconv"

func largestPalindromic(num string) string {
	cnt := [10]int{}
	for _, c := range num {
		cnt[c-'0']++
	}
	n := 0
	maxSingleNum := -1
	if cnt[0]%2 == 1 {
		maxSingleNum = 0
	}
	for i := 1; i < 10; i++ {
		if cnt[i]%2 == 1 && i > maxSingleNum {
			maxSingleNum = i
		}
		n += cnt[i] / 2 * 2
	}
	if n > 0 {
		n += cnt[0] / 2 * 2
	} else {
		if maxSingleNum == -1 {
			return "0"
		}
		return strconv.Itoa(maxSingleNum)
	}
	if maxSingleNum != -1 {
		n++
	}
	res := make([]byte, n)
	k := 9
	for i := 0; i < n/2; i++ {
		if cnt[k] > 1 {
			cnt[k] -= 2
			res[i] = byte('0' + k)
			res[n-i-1] = byte('0' + k)
		} else {
			i--
			k--
		}
	}
	if maxSingleNum != -1 {
		res[n/2] = byte('0' + maxSingleNum)
	}
	return string(res)
}

func main() {
	println(largestPalindromic("444947137"))
}
