/**
 * @author: baowj
 * @time: 2022/9/30 09:39:10
 */
package main

func setZeroes(matrix [][]int) {
	rowSet := map[int]bool{}
	columnSet := map[int]bool{}
	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix[i]); j++ {
			if matrix[i][j] == 0 {
				rowSet[i] = true
				columnSet[j] = true
			}
		}
	}
	for k, _ := range rowSet {
		for j := 0; j < len(matrix[k]); j++ {
			matrix[k][j] = 0
		}
	}
	for k, _ := range columnSet {
		for j := 0; j < len(matrix); j++ {
			matrix[j][k] = 0
		}
	}
}
