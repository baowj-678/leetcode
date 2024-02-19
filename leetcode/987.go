package main

import (
	"sort"
)

func verticalTraversal(root *TreeNode) [][]int {
	type pv struct {
		val int
		row int
	}
	grid := map[int][]pv{}
	var dfs func(node *TreeNode, row, col int)

	dfs = func(node *TreeNode, row, col int) {
		if node == nil {
			return
		}
		if _, ok := grid[col]; !ok {
			grid[col] = []pv{}
		}
		grid[col] = append(grid[col], pv{node.Val, row})

		dfs(node.Left, row+1, col-1)
		dfs(node.Right, row+1, col+1)
	}
	dfs(root, 0, 0)

	type pvs struct {
		vals []pv
		col  int
	}
	pvss := []pvs{}
	for k, v := range grid {
		pvss = append(pvss, pvs{v, k})
	}
	sort.Slice(pvss, func(i, j int) bool {
		return pvss[i].col < pvss[j].col
	})
	res := make([][]int, 0, len(grid))
	for _, s := range pvss {
		tmp := make([]int, 0, len(s.vals))
		sort.Slice(s.vals, func(i, j int) bool {
			if s.vals[i].row == s.vals[j].row {
				return s.vals[i].val < s.vals[j].val
			}
			return s.vals[i].row < s.vals[j].row
		})
		for _, x := range s.vals {
			tmp = append(tmp, x.val)
		}
		res = append(res, tmp)
	}
	return res
}
