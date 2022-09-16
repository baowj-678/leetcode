/**
 * @author: baowj
 * @time: 2022/8/28 22:14:22
 */
package main

func buildMatrix(k int, rowConditions [][]int, colConditions [][]int) [][]int {
	res := make([][]int, k)
	for i := 0; i < k; i++ {
		res[i] = make([]int, k)
	}
	row := topoSort(rowConditions, k)
	if len(row) < k {
		return nil
	}
	col := topoSort(colConditions, k)
	if len(col) < k {
		return nil
	}
	for r, x := range row {
		for c, i := range col {
			if i == x {
				res[r][c] = x
				break
			}
		}
	}
	return res
}

func topoSort(g [][]int, k int) []int {
	d := make([][]int, k+1)
	cnt := make([]int, k+1)
	for _, t := range g {
		if d[t[0]] == nil {
			d[t[0]] = []int{t[1]}
		} else {
			d[t[0]] = append(d[t[0]], t[1])
		}
		cnt[t[1]]++
	}
	res := []int{}
	for {
		has0 := false
		for i := 1; i <= k; i++ {
			if cnt[i] == 0 {
				has0 = true
				res = append(res, i)
				cnt[i] = -1
				for _, x := range d[i] {
					cnt[x]--
				}
			}
		}
		if !has0 {
			break
		}
	}
	return res
}
