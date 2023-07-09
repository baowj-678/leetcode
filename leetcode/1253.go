/**
 * @author: baowj
 * @time: 2023/6/29 10:04
 */
package main

func reconstructMatrix(upper int, lower int, colsum []int) [][]int {
	res := make([][]int, 2)
	res[0] = make([]int, len(colsum))
	res[1] = make([]int, len(colsum))
	count1 := 0
	for i, x := range colsum {
		if x == 2 {
			upper--
			lower--
			res[0][i] = 1
			res[1][i] = 1
		} else if x == 1 {
			count1++
		}
	}
	if upper < 0 || lower < 0 || upper+lower != count1 {
		return [][]int{}
	}
	for i, x := range colsum {
		if x == 1 && upper > 0 {
			res[0][i] = 1
			upper--
		} else if x == 1 {
			res[1][i] = 1
			lower--
		}
	}
	return res
}
