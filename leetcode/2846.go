package main

func minOperationsQueries(n int, edges [][]int, queries [][]int) []int {
	freq := make([][27]int, n)
	neighbors := make([][][]int, n)
	for _, edge := range edges {
		neighbors[edge[0]] = append(neighbors[edge[0]], edge[1:])
		neighbors[edge[1]] = append(neighbors[edge[1]], []int{edge[0], edge[2]})
	}
	flags := make([]bool, n)
	lca := make([][]int, n)
	lca[0] = []int{}
	var dfs func(idx int)
	dfs = func(idx int) {
		if flags[idx] {
			return
		}
		for _, neighbor := range neighbors[idx] {
			if flags[neighbor[0]] {
				continue
			}
			freq[neighbor[0]] = freq[idx]
			freq[neighbor[0]][neighbor[1]]++
			flags[neighbor[0]] = true
			dfs(neighbor[0])
		}
	}
	dfs(0)

	findParent := func(a, b int) int {
		// todo
		return 0
	}

	res := make([]int, 0, len(queries))
	for _, query := range queries {
		f := freq[query[0]]
		for i := 0; i < 27; i++ {
			f[i] += freq[query[1]][i]
		}
		p := findParent(query[0], query[1])

		max, total := 0, 0
		for i := 0; i < 27; i++ {
			f[i] -= freq[p][i] * 2
			total += f[i]
			if f[i] > max {
				max = f[i]
			}
		}

		res = append(res, total-max)
	}
	return res
}
