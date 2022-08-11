/**
 * @author: baowj
 * @time: 2022/8/7 17:32:11
 */
package main

import "fmt"

func reachableNodes(n int, edges [][]int, restricted []int) int {
	list := make([][]int, n)
	flag := make([]bool, n)
	restrictedMap := map[int]bool{}
	for _, edge := range edges {
		list[edge[0]] = append(list[edge[0]], edge[1])
		list[edge[1]] = append(list[edge[1]], edge[0])
	}
	for i := 0; i < len(flag); i++ {
		flag[i] = false
	}
	for _, x := range restricted {
		restrictedMap[x] = true
	}
	return dfs(list, flag, restrictedMap, 0)
}

func dfs(list [][]int, flag []bool, restricted map[int]bool, x int) int {
	if flag[x] {
		return 0
	}
	flag[x] = true
	if _, ok := restricted[x]; ok {
		return 0
	}
	cnt := 1
	for _, y := range list[x] {
		cnt += dfs(list, flag, restricted, y)
	}
	return cnt
}
func main() {
	list := make([][]int, 3)
	list[0] = append(list[0], 9)
	fmt.Println(list)
}
