/**
 * @author: baowj
 * @time: 2022/10/3 16:20:18
 */
package main

func commonFactors(a int, b int) int {
	res := 0
	for i := 1; i <= a && i <= b; i++ {
		if a%i == 0 && b%i == 0 {
			res++
		}
	}
	return res
}
