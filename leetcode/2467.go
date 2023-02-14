/**
 * @author: baowj
 * @time: 2022/11/18 17:21
 */
package main

import (
	"math"
)

func mostProfitablePath(edges [][]int, bob int, amount []int) int {
	next := make([][]int, len(amount))
	for _, edge := range edges {
		if next[edge[0]] == nil {
			next[edge[0]] = []int{edge[1]}
		} else {
			next[edge[0]] = append(next[edge[0]], edge[1])
		}

		if next[edge[1]] == nil {
			next[edge[1]] = []int{edge[0]}
		} else {
			next[edge[1]] = append(next[edge[1]], edge[0])
		}
	}
	preDfs2467(next, amount, 0, -1, bob, 0)
	return dfs2467(next, amount, 0, 0, -1, bob)
}

func dfs2467(next [][]int, amount []int, idx int, sum int, last int, bob int) int {
	sum += amount[idx]
	if idx != 0 && len(next[idx]) == 1 {
		return sum
	}

	max := math.MinInt
	for _, x := range next[idx] {
		if x == last {
			continue
		}
		tmp := dfs2467(next, amount, x, sum, idx, bob)
		if tmp > max {
			max = tmp
		}
	}
	return max
}

func preDfs2467(next [][]int, amount []int, idx int, last int, bob int, cnt int) int {
	cnt++
	if idx == bob {
		amount[idx] = 0
		return cnt
	}
	max := 0
	for _, x := range next[idx] {
		if x == last {
			continue
		}
		tmp := preDfs2467(next, amount, x, idx, bob, cnt)
		if tmp > 0 {
			max = tmp
			if max%2 == 0 {
				if cnt > max/2 {
					amount[idx] = 0
				}
			} else {
				if cnt == max/2+1 {
					amount[idx] /= 2
				} else if cnt > max/2+1 {
					amount[idx] = 0
				}
			}
			break
		}
	}
	return max
}

func main() {
	edges := [][]int{{0, 1}, {1, 2}, {1, 3}, {3, 4}}
	amount := []int{-2, 4, 2, -4, 6}
	print(mostProfitablePath(edges, 3, amount))
}
