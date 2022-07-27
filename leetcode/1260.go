/**
 * @author: baowj
 * @time: 2022/7/20 09:54:28
 */
package main

func shiftGrid(grid [][]int, k int) [][]int {
	width := len(grid[0])
	height := len(grid)
	sum := width * height
	flag := make([]uint64, height)
	for i := 0; i < len(flag); i++ {
		flag[i] = 0
	}
	for index, val, i := 0, -1001, 0; ; {
		if val == -1001 {
			if i >= sum {
				break
			}
			x, y := getXY(i, width)
			if flag[x]&(1<<y) == 0 {
				grid[x][y], val = val, grid[x][y]
				index = (i + k) % sum
			}
			i += 1
		} else {
			x, y := getXY(index, width)
			if flag[x]&(1<<y) == 0 {
				grid[x][y], val = val, grid[x][y]
				flag[x] = flag[x] | (1 << y)
				index = (index + k) % sum
			}
		}
	}
	return grid
}

func getXY(index int, width int) (int, int) {
	x := index / width
	y := index % width
	return x, y
}

func main() {
	a := [][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}
	shiftGrid(a, 1)
}
