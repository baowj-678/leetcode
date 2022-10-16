/**
 * @author: baowj
 * @time: 2022/9/18 13:30:45
 */
package main

func smallestEvenMultiple(n int) int {
	if n%2 == 1 {
		return 2 * n
	}
	return n
}
