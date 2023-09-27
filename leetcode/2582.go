/**
 * @author: baowj
 * @time: 2023/9/26 17:06
 */
package main

func passThePillow(n int, time int) int {
	direction := time/(n-1) + 1
	time %= n - 1
	if direction%2 == 1 {
		return 1 + time
	} else {
		return n - time
	}
}
