/**
 * @author: baowj
 * @time: 2022/11/26 21:53:12
 * @email: bwj_678@qq.com
 */
package main

import (
	"math"
	"strconv"
)

func reachableNodes(edges [][]int, maxMoves int, n int) int {
	moves := make([]int, n)
	closestPath(edges, moves, n)
	m := map[string]int{}
	for _, edge := range edges {
		s := strconv.Itoa(edge[0]) + "-" + strconv.Itoa(edge[1])
		s_ := strconv.Itoa(edge[1]) + "-" + strconv.Itoa(edge[0])
		mv, mv_ := 0, 0
		if v, ok := m[s]; ok {
			mv = v
		}
		if v, ok := m[s_]; ok {
			mv_ = v
		}
		if moves[edge[0]] < maxMoves {
			gain := maxMoves - moves[edge[0]]
			if gain > edge[2] {
				gain = edge[2]
			}
			left := edge[2] - mv_
			if gain > left {
				gain = left
			}
			if gain > mv {
				m[s] = gain
				mv = gain
			}
		}

		if moves[edge[1]] < maxMoves {
			gain := maxMoves - moves[edge[1]]
			if gain > edge[2] {
				gain = edge[2]
			}
			left := edge[2] - mv
			if gain > left {
				gain = left
			}
			if gain > mv_ {
				m[s_] = gain
			}
		}
	}

	res := 0
	for _, move := range moves {
		if move <= maxMoves {
			res++
		}
	}
	for _, v := range m {
		res += v
	}
	return res
}

func closestPath(edges [][]int, moves []int, n int) {
	for i := 0; i < n; i++ {
		moves[i] = math.MaxInt
	}
	moves[0] = 0
	for i := 0; i < n-1; i++ {
		for _, edge := range edges {
			if moves[edge[0]] != math.MaxInt {
				if moves[edge[0]]+edge[2]+1 < moves[edge[1]] {
					moves[edge[1]] = moves[edge[0]] + edge[2] + 1
				}
			}

			if moves[edge[1]] != math.MaxInt {
				if moves[edge[1]]+edge[2]+1 < moves[edge[0]] {
					moves[edge[0]] = moves[edge[1]] + edge[2] + 1
				}
			}
		}
	}
}

func main() {
	edges := [][]int{{0, 1, 10}, {0, 2, 1}, {1, 2, 2}}
	println(reachableNodes(edges, 6, 3))
}
