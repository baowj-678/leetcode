package main

import "sort"

func minimumOperations6235(root *TreeNode) int {
	res := 0
	queue := []*TreeNode{root, nil}
	vec := []int{}
	for len(queue) > 1 {
		head := queue[0]
		queue = queue[1:]
		if head == nil {
			// do something
			res += f6235(vec)
			queue = append(queue, nil)
			vec = []int{}
			continue
		}
		if head.Left != nil {
			queue = append(queue, head.Left)
		}
		if head.Right != nil {
			queue = append(queue, head.Right)
		}
		vec = append(vec, head.Val)
	}
	if len(vec) != 0 {
		// do something
		res += f6235(vec)
	}
	return res
}

func union6235(x, y int, m map[int]int) {
	xP := find6235(x, m)
	yP := find6235(y, m)
	m[xP] = yP
}

func find6235(x int, m map[int]int) int {
	start := x
	for m[x] != x {
		x = m[x]
	}
	for start != x {
		m[start], start = x, m[start]
	}
	return x
}

func f6235(vec []int) int {
	sorted := make([]int, len(vec))
	copy(sorted, vec)
	sort.Ints(sorted)
	m := map[int]int{}
	for _, x := range vec {
		m[x] = x
	}
	for i := 0; i < len(vec); i++ {
		union6235(sorted[i], vec[i], m)
	}
	set := map[int]int{}
	for _, x := range vec {
		set[find6235(x, m)] = 1
	}
	return len(vec) - len(set)
}
