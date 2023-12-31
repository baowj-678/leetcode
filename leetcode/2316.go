package main

func countPairs2316(n int, edges [][]int) int64 {
	parent := make([]int, n)
	for i := range parent {
		parent[i] = i
	}
	find := func(x int) int {
		tmp := x
		for parent[x] != x {
			x = parent[x]
		}
		for parent[tmp] != x {
			parent[tmp], tmp = x, parent[tmp]
		}
		return x
	}

	union := func(x, y int) {
		x = find(x)
		y = find(y)
		parent[x] = y
	}

	for _, edge := range edges {
		union(edge[0], edge[1])
	}

	m := map[int]int{}
	for i := range parent {
		m[find(i)]++
	}

	groups := make([]int, 0, len(m))
	for _, v := range m {
		groups = append(groups, v)
	}
	sums := make([]int, len(m))
	s := 0
	for i := len(groups) - 1; i >= 0; i-- {
		sums[i] = s
		s += groups[i]
	}

	res := int64(0)
	for i, g1 := range groups {
		res += int64(g1) * int64(sums[i])
	}

	return res
}
