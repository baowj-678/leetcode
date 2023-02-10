/**
 * @author: baowj
 * @time: 2023/2/9 22:08:28
 * @email: bwj_678@qq.com
 */
package main

import "sort"

func matrixRankTransform(matrix [][]int) [][]int {
	m, n := len(matrix), len(matrix[0])
	next := make([]map[int]struct{}, m*n)
	// uf
	hasLast := make([]bool, m*n)
	uf := make([]int, n*m)
	for i := 0; i < len(uf); i++ {
		uf[i] = i
		next[i] = map[int]struct{}{}
	}
	find := func(uf []int, x int) int {
		tmp := x
		for uf[tmp] != tmp {
			tmp = uf[tmp]
		}
		for uf[x] != tmp {
			x, uf[x] = uf[x], tmp
		}
		return tmp
	}

	union := func(uf []int, x, y int) {
		xp, yp := find(uf, x), find(uf, y)
		uf[yp] = xp
	}

	// row init
	for i, row := range matrix {
		tmp := make([][2]int, len(row))
		for j, x := range row {
			tmp[j] = [2]int{x, i*n + j}
		}
		sort.Slice(tmp, func(i, j int) bool {
			return tmp[i][0] < tmp[j][0]
		})
		for k, x := range tmp {
			if k+1 < len(tmp) {
				if x[0] < tmp[k+1][0] {
					next[tmp[k][1]][tmp[k+1][1]] = struct{}{}
				} else {
					union(uf, x[1], tmp[k+1][1])
				}
			}
		}
	}

	// column init
	for j := 0; j < n; j++ {
		tmp := make([][2]int, m)
		for i := 0; i < m; i++ {
			tmp[i] = [2]int{matrix[i][j], i*n + j}
		}
		sort.Slice(tmp, func(i, j int) bool {
			return tmp[i][0] < tmp[j][0]
		})
		for k, x := range tmp {
			if k+1 < len(tmp) {
				if x[0] < tmp[k+1][0] {
					next[tmp[k][1]][tmp[k+1][1]] = struct{}{}
				} else {
					union(uf, x[1], tmp[k+1][1])
				}
			}
		}
	}

	// union
	for i := 0; i < n*m; i++ {
		ip := find(uf, i)
		if ip != i {
			if next[i] != nil {
				for key, _ := range next[i] {
					kp := find(uf, key)
					hasLast[kp] = true
					next[ip][kp] = struct{}{}
				}
			}
			next[i] = nil
		} else {
			if next[i] != nil {
				for k, _ := range next[i] {
					kp := find(uf, k)
					if k != kp {
						delete(next[i], k)
					}
					hasLast[kp] = true
					next[i][kp] = struct{}{}
				}
			}
		}
	}

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			matrix[i][j] = 1
		}
	}
	// toposort

	for i := 0; i < n*m; i++ {
		if !hasLast[i] {
			topoSort1632(next, i, 1, matrix, n)
		}
	}

	for i := 0; i < n*m; i++ {
		ip := find(uf, i)
		if i != ip {
			matrix[i/n][i%n] = matrix[ip/n][ip%n]
		}
	}
	return matrix
}

func topoSort1632(next []map[int]struct{}, idx int, v int, matrix [][]int, n int) {
	if v > matrix[idx/n][idx%n] {
		matrix[idx/n][idx%n] = v
	}
	if next[idx] == nil {
		return
	}
	for key, _ := range next[idx] {
		topoSort1632(next, key, v+1, matrix, n)
	}
}
