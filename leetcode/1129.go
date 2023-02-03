/**
 * @author: baowj
 * @time: 2023/2/3 09:56:42
 * @email: bwj_678@qq.com
 */
package main

func shortestAlternatingPaths(n int, redEdges [][]int, blueEdges [][]int) []int {
	redPath, bluePath := make([]int, n), make([]int, n)
	redNext, blueNext := make([][]int, n), make([][]int, n)
	res := make([]int, n)
	for i := 1; i < n; i++ {
		redPath[i], bluePath[i] = 1e5, 1e5
	}
	for _, edge := range redEdges {
		redNext[edge[0]] = append(redNext[edge[0]], edge[1])
	}

	for _, edge := range blueEdges {
		blueNext[edge[0]] = append(blueNext[edge[0]], edge[1])
	}

	maxIter := len(redEdges) + len(blueEdges)
	for i := 0; i < maxIter; i++ {
		for s := 0; s < n; s++ {
			tmp := redPath[s]
			for _, next := range redNext[s] {
				if bluePath[next] > tmp+1 {
					bluePath[next] = tmp + 1
				}
			}
			tmp = bluePath[s]
			for _, next := range blueNext[s] {
				if redPath[next] > tmp+1 {
					redPath[next] = tmp + 1
				}
			}
		}
	}
	for i := 0; i < n; i++ {
		if redPath[i] < bluePath[i] {
			res[i] = redPath[i]
		} else {
			res[i] = bluePath[i]
		}
		if res[i] == 1e5 {
			res[i] = -1
		}
	}
	return res
}
