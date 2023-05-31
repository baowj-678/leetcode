/**
 * @author: baowj
 * @time: 2023/5/24 9:58
 */
package main

func frogPosition(n int, edges [][]int, t int, target int) float64 {
	next := make([]map[int]struct{}, 101)
	for _, edge := range edges {
		if next[edge[0]] == nil {
			next[edge[0]] = map[int]struct{}{}
		}
		if next[edge[1]] == nil {
			next[edge[1]] = map[int]struct{}{}
		}
		next[edge[0]][edge[1]] = struct{}{}
		next[edge[1]][edge[0]] = struct{}{}
	}
	dfs1377(next, 1, 0)
	res := dfs13772(next, 1, 1, target, t)
	if res < 0 {
		res = 0
	}
	return res
}

func dfs1377(next []map[int]struct{}, n int, p int) {
	delete(next[n], p)
	for k, _ := range next[n] {
		dfs1377(next, k, n)
	}
}

func dfs13772(next []map[int]struct{}, n int, v float64, target int, t int) float64 {
	if t < 0 {
		return -1
	}
	if n == target {
		if len(next[target]) == 0 {
			return v
		} else {
			if t == 0 {
				return v
			} else {
				return -1
			}
		}
	}
	v /= float64(len(next[n]))
	for i, _ := range next[n] {
		tmp := dfs13772(next, i, v, target, t-1)
		if tmp > 0 {
			return tmp
		}
	}
	return -1
}
