/**
 * @author: baowj
 * @time: 2023/6/26 9:57
 */
package main

func pivotInteger(n int) int {
	check := func(x int) bool {
		if (1+x)*x/2 == (x+n)*(n-x+1)/2 {
			return true
		}
		return false
	}
	for i := 1; i <= n; i++ {
		if check(i) {
			return i
		}
	}
	return -1
}
