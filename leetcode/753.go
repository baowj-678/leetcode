/**
 * @author: baowj
 * @time: 2023/1/10 11:05:14
 * @email: bwj_678@qq.com
 */
package main

import "strconv"

func crackSafe(n int, k int) string {
	sum := 1
	for i := 0; i < n; i++ {
		sum *= k
	}
	graph := make([][]int, sum)
	for i := 0; i < sum; i++ {
		graph[i] = getNextN(i, k, sum)
	}
	s, _ := dfs753(0, 0, sum, graph, make([]bool, sum))
	for i := 0; i < n; i++ {
		s = "0" + s
	}
	return s
}

func dfs753(x int, cnt int, sum int, graph [][]int, visited []bool) (string, bool) {
	if visited[x] {
		return "", false
	}
	visited[x] = true
	cnt++
	if cnt == sum {
		return "", true
	}
	for i, next := range graph[x] {
		if s, ok := dfs753(next, cnt, sum, graph, visited); ok {
			return strconv.Itoa(i) + s, true
		}
	}
	visited[x] = false
	return "", false
}

func getNextN(x int, k int, max int) []int {
	x = (x * k) % max
	res := make([]int, k)
	for i := 0; i < k; i++ {
		res[i] = x + i
	}
	return res
}

func main() {
	println(crackSafe(1, 2))
}
