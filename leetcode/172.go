/**
 * @author: baowj
 * @time: 2022/8/28 21:01:02
 */
package main

func trailingZeroes(n int) int {
	ans := 0
	for i := 1; ; i++ {
		tmp := n / pow5(i)
		if tmp == 0 {
			break
		}
		ans += tmp
	}
	return ans
}

func pow5(p int) int {
	res := 1
	for i := 0; i < p; i++ {
		res *= 5
	}
	return res
}
