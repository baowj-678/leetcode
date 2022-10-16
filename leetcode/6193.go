/**
 * @author: baowj
 * @time: 2022/10/3 16:23:50
 */
package main

func maxSum(grid [][]int) int {
	res := 0
	for i := 0; i+2 < len(grid); i++ {
		for j := 0; j+2 < len(grid[0]); j++ {
			tmp := grid[i][j] + grid[i][j+1] + grid[i][j+2] + grid[i+1][j+1] + grid[i+2][j] + grid[i+2][j+1] + grid[i+2][j+2]
			if tmp > res {
				res = tmp
			}
		}
	}
	return res
}
