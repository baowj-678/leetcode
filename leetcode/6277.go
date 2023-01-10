/**
 * @author: baowj
 * @time: 2022/11/27 15:55:34
 * @email: bwj_678@qq.com
 */
package main

func onesMinusZeros(grid [][]int) [][]int {
	oneRows := make([]int, len(grid))
	oneCols := make([]int, len(grid[0]))
	zeroRows := make([]int, len(grid))
	zeroCols := make([]int, len(grid[0]))
	for i := 0; i < len(grid); i++ {
		zero, one := 0, 0
		for j := 0; j < len(grid[i]); j++ {
			if grid[i][j] == 0 {
				zero++
			} else {
				one++
			}
		}
		oneRows[i] = one
		zeroRows[i] = zero
	}

	for j := 0; j < len(grid[0]); j++ {
		zero, one := 0, 0
		for i := 0; i < len(grid); i++ {
			if grid[i][j] == 0 {
				zero++
			} else {
				one++
			}
		}
		oneCols[j] = one
		zeroCols[j] = zero
	}
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[i]); j++ {
			grid[i][j] = oneRows[i] + oneCols[j] - zeroRows[i] - zeroCols[i]
		}
	}
	return grid
}
