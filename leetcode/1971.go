/**
 * @author: baowj
 * @time: 2022/12/19 21:44:18
 * @email: bwj_678@qq.com
 */
package main

func validPath(n int, edges [][]int, source int, destination int) bool {
	parent := make([]int, n)
	for i := 0; i < n; i++ {
		parent[i] = i
	}
	for _, edge := range edges {
		union1971(parent, edge[0], edge[1])
	}
	return find1971(parent, source) == find1971(parent, destination)
}

func union1971(parent []int, x, y int) {
	xP, yP := find1971(parent, x), find1971(parent, y)
	parent[xP] = yP
}

func find1971(parent []int, x int) int {
	tmp := x
	for x != parent[x] {
		x = parent[x]
	}
	for tmp != x {
		tmp, parent[tmp] = parent[tmp], x
	}
	return x
}
