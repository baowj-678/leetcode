/**
 * @author: baowj
 * @time: 2022/9/22 09:39:27
 */
package main

func canFormArray(arr []int, pieces [][]int) bool {
	for i := 0; i < len(arr); {
		tmp := i
		for j := 0; j < len(pieces); j++ {
			if i < len(arr) && arr[i] == pieces[j][0] {
				for k := 0; k < len(pieces[j]); k++ {
					if arr[i] != pieces[j][k] {
						return false
					}
					i++
				}
			}
		}
		if tmp == i {
			return false
		}
	}
	return true
}
