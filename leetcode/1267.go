package main

func countServers(grid [][]int) int {
	res := map[int]struct{}{}
	w := len(grid[0])
	h := len(grid)
	for i := 0; i < h; i++ {
		row := grid[i]
		cnt := 0
		for j := 0; j < len(row); j++ {
			if row[j] == 1 {
				cnt++
			}
			if cnt > 1 {
				break
			}
		}

		if cnt > 1 {
			for j := 0; j < len(row); j++ {
				if row[j] == 1 {
					res[i*w+j] = struct{}{}
				}
			}
		}
	}

	for j := 0; j < w; j++ {
		cnt := 0
		for i := 0; i < h; i++ {
			if grid[i][j] == 1 {
				cnt++
			}
			if cnt > 1 {
				break
			}
		}

		if cnt > 1 {
			for i := 0; i < h; i++ {
				if grid[i][j] == 1 {
					res[i*w+j] = struct{}{}
				}
			}
		}
	}
	return len(res)
}
