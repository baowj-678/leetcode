/**
 * @author: baowj
 * @time: 2023/5/29 20:43
 */
package main

import "math"

func shortestPathBinaryMatrix(grid [][]int) int {
	if grid[0][0] == 1 {
		return -1
	}
	q := []int{0}
	w, h := len(grid[0]), len(grid)
	minPath := make([][]int, h)
	for i := 0; i < h; i++ {
		minPath[i] = make([]int, w)
		for j := 0; j < w; j++ {
			minPath[i][j] = math.MaxInt
		}
	}
	minPath[0][0] = 1
	for len(q) > 0 {
		v := q[0]
		q = q[1:]
		w_, h_ := v%w, v/w
		p := minPath[h_][w_]
		// h_, w_+1
		if w_+1 < w && grid[h_][w_+1] == 0 {
			if p+1 < minPath[h_][w_+1] {
				minPath[h_][w_+1] = p + 1
				q = append(q, h_*w+w_+1)
			}
		}

		// h_, w_-1
		if w_-1 >= 0 && grid[h_][w_-1] == 0 {
			if p+1 < minPath[h_][w_-1] {
				minPath[h_][w_-1] = p + 1
				q = append(q, h_*w+w_-1)
			}
		}

		// h_+1, w_
		if h_+1 < h && grid[h_+1][w_] == 0 {
			if p+1 < minPath[h_+1][w_] {
				minPath[h_+1][w_] = p + 1
				q = append(q, (h_+1)*w+w_)
			}
		}

		// h_-1, w_
		if h_-1 >= 0 && grid[h_-1][w_] == 0 {
			if p+1 < minPath[h_-1][w_] {
				minPath[h_-1][w_] = p + 1
				q = append(q, (h_-1)*w+w_)
			}
		}

		// h_-1, w_-1
		if h_-1 >= 0 && w_-1 >= 0 && grid[h_-1][w_-1] == 0 {
			if p+1 < minPath[h_-1][w_-1] {
				minPath[h_-1][w_-1] = p + 1
				q = append(q, (h_-1)*w+w_-1)
			}
		}

		// h_-1, w_+1
		if h_-1 >= 0 && w_+1 < w && grid[h_-1][w_+1] == 0 {
			if p+1 < minPath[h_-1][w_+1] {
				minPath[h_-1][w_+1] = p + 1
				q = append(q, (h_-1)*w+w_+1)
			}
		}

		// h_+1, w_+1
		if h_+1 < h && w_+1 < w && grid[h_+1][w_+1] == 0 {
			if p+1 < minPath[h_+1][w_+1] {
				minPath[h_+1][w_+1] = p + 1
				q = append(q, (h_+1)*w+w_+1)
			}
		}

		// h_+1, w_-1
		if h_+1 < h && w_-1 >= 0 && grid[h_+1][w_-1] == 0 {
			if p+1 < minPath[h_+1][w_-1] {
				minPath[h_+1][w_-1] = p + 1
				q = append(q, (h_+1)*w+w_-1)
			}
		}
	}
	if minPath[h-1][w-1] == math.MaxInt {
		return -1
	}
	return minPath[h-1][w-1]
}
