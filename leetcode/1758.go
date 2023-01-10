/**
 * @author: baowj
 * @time: 2022/12/3 11:28:10
 * @email: bwj_678@qq.com
 */
package main

func minOperations1758(s string) int {
	cnt0, cnt1 := 0, 0
	cur0, cur1 := 0, 1
	for _, c := range s {
		c -= '0'
		if cur0 != int(c) {
			cnt0++
		}
		cur0 = f1758(cur0)

		if cur1 != int(c) {
			cnt1++
		}
		cur1 = f1758(cur1)
	}
	if cnt0 > cnt1 {
		return cnt1
	}
	return cnt0
}

func f1758(x int) int {
	if x == 0 {
		return 1
	} else {
		return 0
	}
}
