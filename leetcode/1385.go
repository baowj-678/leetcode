/**
 * @author: baowj
 * @time: 2022/9/8 10:14:04
 */
package main

func findTheDistanceValue(arr1 []int, arr2 []int, d int) int {
	res := 0
	for _, x := range arr1 {
		for _, y := range arr2 {
			if abs1385(x-y) <= d {
				res--
				break
			}
		}
		res++
	}
	return res
}

func abs1385(x int) int {
	if x < 0 {
		x = -x
	}
	return x
}
