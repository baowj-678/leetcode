/**
 * @author: baowj
 * @time: 2022/8/28 23:18:49
 */
package main

func checkIfExist(arr []int) bool {
	for _, x := range arr {
		for _, y := range arr {
			if x == 2*y {
				return true
			}
		}
	}
	return false
}
