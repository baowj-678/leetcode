/**
 * @author: baowj
 * @time: 2022/12/10 22:05:22
 * @email: bwj_678@qq.com
 */
package main

func checkPowersOfThree(n int) bool {
	for n > 0 {
		if n%3 == 0 {
			n /= 3
		} else if n%3 == 1 {
			n -= 1
		} else if n%3 == 2 {
			return false
		}
	}
	return true
}
