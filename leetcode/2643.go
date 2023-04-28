/**
 * @author: baowj
 * @time: 2023/4/17 10:26
 */
package main

func rowAndMaximumOnes(mat [][]int) []int {
	res := make([]int, 2)
	for i, arr := range mat {
		cnt := 0
		for _, x := range arr {
			if x == 1 {
				cnt++
			}
		}
		if cnt > res[1] {
			res = []int{i, cnt}
		}
	}
	return res
}
