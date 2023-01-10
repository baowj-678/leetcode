/**
 * @author: baowj
 * @time: 2023/1/8 11:35:37
 * @email: bwj_678@qq.com
 */
package main

func countEven(num int) int {
	cnt := 0
	for i := 1; i <= num; i++ {
		if isEven2180(i) {
			cnt++
		}
	}
	return cnt
}

func isEven2180(num int) bool {
	s := 0
	for num > 0 {
		s += num % 10
		num /= 10
	}
	return s%2 == 0
}
