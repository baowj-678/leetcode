/**
 * @author: baowj
 * @time: 2022/12/15 22:12:11
 * @email: bwj_678@qq.com
 */
package main

func getLucky(s string, k int) int {
	res := 0
	for _, c := range s {
		num := c - 'a' + 1
		for num > 0 {
			res += int(num % 10)
			num /= 10
		}
	}
	for i := 1; i < k; i++ {
		tmp := res
		res = 0
		for tmp > 0 {
			res += tmp % 10
			tmp /= 10
		}
	}
	return res
}
