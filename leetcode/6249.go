/**
 * @author: baowj
 * @time: 2022/11/27 15:45:48
 * @email: bwj_678@qq.com
 */
package main

func numberOfCuts(n int) int {
	if n == 1 {
		return 0
	}
	if n%2 == 0 {
		return n / 2
	} else {
		return n
	}
}
