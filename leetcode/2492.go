/**
 * @author: baowj
 * @time: 2022/12/5 23:51:03
 * @email: bwj_678@qq.com
 */
package main

import "math"

func minScore(n int, roads [][]int) int {
	parent, dis := make([]int, n+1), make([]int, n+1)
	for i := 0; i < len(parent); i++ {
		parent[i] = i
		dis[i] = math.MaxInt
	}
	for _, road := range roads {
		union2492(parent, dis, road[0], road[1], road[2])
	}

}

func find2492(parent []int, x int) int {
	tmp := x
	for parent[x] != x {
		x = parent[x]
	}
	for tmp != x {
		tmp, parent[tmp] = parent[tmp], x
	}
	return x
}

func union2492(parent, dis []int, x, y, d int) {
	xP, yP := find2492(parent, x), find2492(parent, y)
	parent[yP] = xP
	if dis[yP] < dis[xP] {
		dis[xP] = dis[yP]
	}
	if dis[xP] > d {
		dis[xP] = d
	}
}
