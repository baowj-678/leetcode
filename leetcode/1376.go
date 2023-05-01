/**
 * @author: baowj
 * @time: 2023/5/1 10:11:48
 * @email: bwj_678@qq.com
 */
package main

func numOfMinutes(n int, headID int, manager []int, informTime []int) int {
	nexts := make([][]int, n)
	for i, m := range manager {
		if m == -1 {
			continue
		}
		nexts[m] = append(nexts[m], i)
	}
	return dfs1376(nexts, informTime, headID)
}

func dfs1376(nexts [][]int, informTime []int, idx int) int {
	res := 0
	max := func(x, y int) int {
		if x < y {
			x = y
		}
		return x
	}
	for _, next := range nexts[idx] {
		res = max(res, dfs1376(nexts, informTime, next))
	}
	return res + informTime[idx]
}
