/**
 * @author: baowj
 * @time: 2022/9/25 21:39:41
 */
package main

import "sort"

type UF struct {
	nums []int
}

func Construct(size int) UF {
	nums := make([]int, size)
	for i := 0; i < size; i++ {
		nums[i] = i
	}
	return UF{nums: nums}
}
func (this *UF) Union(x int, y int) {
	x = this.Find(x)
	y = this.Find(y)
	if x != y {
		this.nums[y] = x
	}
}

func (this *UF) Find(x int) int {
	tmp := x
	for tmp != this.nums[tmp] {
		tmp = this.nums[tmp]
	}
	for x != tmp {
		x, this.nums[x] = this.nums[x], tmp
	}
	return tmp
}

func numberOfGoodPaths(vals []int, edges [][]int) int {
	sort.Slice(edges, func(i, j int) bool {
		iMax := vals[edges[i][0]]
		if iMax < vals[edges[i][1]] {
			iMax = vals[edges[i][1]]
		}
		jMax := vals[edges[j][0]]
		if jMax < vals[edges[j][1]] {
			jMax = vals[edges[j][1]]
		}
		return iMax < jMax
	})

	idx2val := make([][2]int, len(vals))
	for i := 0; i < len(vals); i++ {
		idx2val[i] = [2]int{i, vals[i]}
	}
	sort.Slice(idx2val, func(i, j int) bool {
		return idx2val[i][1] < idx2val[j][1]
	})

	res := 0
	uf := Construct(len(vals))
	k := 0
	for i := 0; i < len(idx2val); i++ {
		val := idx2val[i][1]
		for k < len(edges) && vals[edges[k][0]] <= val && vals[edges[k][1]] <= val {
			// add edge
			uf.Union(edges[k][0], edges[k][1])
			k++
		}
		m := map[int]int{}
		for ; i < len(vals) && idx2val[i][1] == val; i++ {
			x := uf.Find(idx2val[i][0])
			if p, ok := m[x]; ok {
				m[x] = p + 1
			} else {
				m[x] = 1
			}
		}

		for _, v := range m {
			res += v * (v - 1) / 2
		}

		if i < len(vals) && idx2val[i][1] != val {
			i--
		}
	}
	res += len(vals)
	return res
}
