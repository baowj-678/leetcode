/**
 * @author: baowj
 * @time: 2022/9/4 16:46:54
 */
package main

import "fmt"

func numSpecial(mat [][]int) int {
	rows := make([]int, len(mat))
	cols := make([]int, len(mat[0]))
	for i := 0; i < len(mat); i++ {
		cnt := 0
		col := 0
		for j := 0; j < len(mat[i]); j++ {
			if mat[i][j] == 1 {
				cnt++
				col = j
			}
		}
		if cnt == 1 {
			rows[i] = col
		} else {
			rows[i] = -1
		}
	}

	for i := 0; i < len(mat[0]); i++ {
		cnt := 0
		row := 0
		for j := 0; j < len(mat); j++ {
			if mat[j][i] == 1 {
				cnt++
				row = j
			}
		}
		if cnt == 1 {
			cols[i] = row
		} else {
			cols[i] = -1
		}
	}
	fmt.Println(rows)
	fmt.Println(cols)
	res := 0
	for i, row := range rows {
		if row != -1 {
			if cols[row] == i {
				res++
			}
		}
	}
	return res
}
