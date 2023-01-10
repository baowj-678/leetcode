/**
 * @author: baowj
 * @time: 2022/11/23 22:25:40
 * @email: bwj_678@qq.com
 */
package main

func minimumFuelCost(roads [][]int, seats int) int64 {
	graph := make([][]int, len(roads)+1)
	for _, road := range roads {
		if graph[road[0]] == nil {
			graph[road[0]] = []int{road[1]}
		} else {
			graph[road[0]] = append(graph[road[0]], road[1])
		}
		if graph[road[1]] == nil {
			graph[road[1]] = []int{road[0]}
		} else {
			graph[road[1]] = append(graph[road[1]], road[0])
		}
	}
	_, _, cost := dfs2477(graph, -1, 0, seats)
	return cost
}

func dfs2477(graph [][]int, lastp, p, seats int) (cars, num int, cost int64) {
	cars, num, cost = 0, 1, 0
	for _, next := range graph[p] {
		if next != lastp {
			cars_, num_, cost_ := dfs2477(graph, p, next, seats)
			num += num_
			cost += cost_ + int64(cars_)
		}
	}
	cars = (num + seats - 1) / seats
	return
}
