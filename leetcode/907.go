/**
 * @author: baowj
 * @time: 2022/11/1 16:29:27
 */
package main

const MAX907 = 10e9 + 7

func sumSubarrayMins(arr []int) int {
	res := 0
	for i, x := range arr {
		l, r := i, i
		for ; l >= 0; l-- {

		}
		res = res + l*x*r%MAX907
	}
	return res
}
