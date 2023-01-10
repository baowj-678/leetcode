/**
 * @author: baowj
 * @time: 2022/12/17 19:57:30
 * @email: bwj_678@qq.com
 */
package main

import (
	"fmt"
	"math"
	"sort"
)

func distanceLimitedPathsExist(n int, edgeList [][]int, queries [][]int) []bool {
	sort.Slice(edgeList, func(i, j int) bool {
		return edgeList[i][2] < edgeList[j][2]
	})
	for i := 0; i < len(queries); i++ {
		queries[i] = append(queries[i], i)
	}
	res := make([]bool, len(queries))
	sort.Slice(queries, func(i, j int) bool {
		return queries[i][2] < queries[j][2]
	})
	uf := make([]int, n+1)
	for i := 0; i <= n; i++ {
		uf[i] = i
	}
	last := 0
	lastEdge := 0
	queries = append(queries, []int{0, 0, math.MaxInt})
	for i := 0; i < len(queries); i++ {
		if queries[i][2] > queries[last][2] {
			limit := queries[last][2]
			for ; lastEdge < len(edgeList) && edgeList[lastEdge][2] < limit; lastEdge++ {
				union1697(uf, edgeList[lastEdge][0], edgeList[lastEdge][1])
			}
			for j := last; j < i; j++ {
				if find1697(uf, queries[j][0]) == find1697(uf, queries[j][1]) {
					res[queries[j][3]] = true
				} else {
					res[queries[j][3]] = false
				}
			}
			last = i
		}
	}
	return res
}

func union1697(uf []int, x, y int) {
	xP, yP := find1697(uf, x), find1697(uf, y)
	uf[xP] = yP
}

func find1697(uf []int, x int) int {
	tmp := x
	for uf[x] != x {
		x = uf[x]
	}
	for x != tmp {
		tmp, uf[tmp] = uf[tmp], x
	}
	return x
}

func main() {
	edgeList := [][]int{{0, 1, 2}, {1, 2, 4}, {2, 0, 8}, {1, 0, 16}}
	queries := [][]int{{0, 1, 2}, {0, 2, 5}}
	fmt.Println(distanceLimitedPathsExist(3, edgeList, queries))
}
