/**
 * @author: baowj
 * @time: 2022/10/27 17:03:32
 */
package main

func shortestBridge(grid [][]int) int {
	parent := make([]int, len(grid)*len(grid[0]))
	width := len(grid[0])
	for i := 0; i < len(parent); i++ {
		parent[i] = i
	}
	startX, startY := -1, -1
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[i]); j++ {
			if grid[i][j] == 1 {
				startX, startY = i, j
				if i > 0 && grid[i-1][j] == 1 {
					union934(parent, i*width+j, (i-1)*width+j)
				}
				if j > 0 && grid[i][j-1] == 1 {
					union934(parent, i*width+j, i*width+j-1)
				}
				if i+1 < len(grid) && grid[i+1][j] == 1 {
					union934(parent, i*width+j, (i+1)*width+j)
				}
				if j+1 < width && grid[i][j+1] == 1 {
					union934(parent, i*width+j, j*width+j+1)
				}
			}
		}
	}
	bfs := []int{}
	res := 0
	for i := 0; i < len(grid); i++ {
		for j := 0; j < width; j++ {
			if find934(parent, i*width+j) == find934(parent, startX*width+startY) {
				bfs = append(bfs, i*width+j)
			}
		}
	}

	bfs = append(bfs, -1)
	for len(bfs) > 1 {
		x := bfs[0]
		bfs = bfs[1:]
		if x == -1 {
			res++
			bfs = append(bfs, -1)
		}
		i, j := x/width, x%width
		if i > 0 {
			if grid[i-1][j] == 0 {
				bfs = append(bfs, (i-1)*width+j)

			}

		}
	}

}

func union934(parent []int, x int, y int) {
	px, py := find934(parent, x), find934(parent, y)
	if px != py {
		parent[px] = py
	}
}

func find934(parent []int, x int) int {
	tmp := x
	for parent[x] != x {
		x = parent[x]
	}
	for parent[tmp] != x {
		parent[tmp], tmp = x, parent[tmp]
	}
	return x
}
