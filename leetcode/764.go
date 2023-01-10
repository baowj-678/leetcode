package main

func orderOfLargestPlusSign(n int, mines [][]int) int {
	graph := make([][]int, n)
	colgraph := make([][]int, n)
	rowgraph := make([][]int, n)
	for i := 0; i < len(graph); i++ {
		graph[i] = make([]int, n)
		colgraph[i] = make([]int, n)
		rowgraph[i] = make([]int, n)
		for j := 0; j < n; j++ {
			graph[i][j] = 1
		}
	}
	for _, mine := range mines {
		graph[mine[0]][mine[1]] = 0
	}

	for i := 0; i < n; i++ {
		left, right := 0, 0
		for j := 0; j < n; j++ {
			if graph[i][j] == 0 {
				left, right = j+1, j+1
				rowgraph[i][j] = 0
			} else {
				for right < n && graph[i][right] == 1 {
					right++
				}
				right--
				length := j - left + 1
				if length > right-j+1 {
					length = right - j + 1
				}
				rowgraph[i][j] = length
			}
		}
	}

	for j := 0; j < n; j++ {
		left, right := 0, 0
		for i := 0; i < n; i++ {
			if graph[i][j] == 0 {
				left, right = i+1, i+1
			} else {
				for right < n && graph[right][j] == 1 {
					right++
				}
				right--
				length := i - left + 1
				if length > right-i+1 {
					length = right - i + 1
				}
				colgraph[i][j] = length
			}
		}
	}
	res := 0
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			max := colgraph[i][j]
			if rowgraph[i][j] < max {
				max = rowgraph[i][j]
			}
			if max > res {
				res = max
			}
		}
	}
	return res
}

func main() {
	println(orderOfLargestPlusSign(5, [][]int{
		{4, 2},
	}))
}
