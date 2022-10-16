/**
 * @author: baowj
 * @time: 2022/10/8 22:05:19
 */
package main

func threeEqualParts(arr []int) []int {
	has1 := false
	for i := 0; i < len(arr); i++ {
		if arr[i] == 1 {
			has1 = true
			break
		}
	}
	if !has1 {
		return []int{0, 2}
	}
	for l, r := 0, 0; l < len(arr); l++ {
		i, j := 0, l+1
		for i <= l && arr[i] == 0 {
			i++
		}
		begin := i
		for j < len(arr) && arr[j] == 0 {
			j++
		}
		for i <= l && j < len(arr) && arr[i] == arr[j] {
			i++
			j++
		}
		r = j
		if i == l+1 {
			for j < len(arr) && arr[j] == 0 {
				j++
			}
			for begin <= l && j < len(arr) {
				if arr[begin] == arr[j] {
					begin++
					j++
				} else {
					break
				}
			}
			if begin == l+1 && j == len(arr) {
				return []int{l, r}
			}
		}
	}
	return []int{-1, -1}
}
