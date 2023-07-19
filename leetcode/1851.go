/**
 * @author: baowj
 * @time: 2023/7/18 10:21
 */
package main

import (
	"container/heap"
	"fmt"
	"sort"
)

type Heap1851 [][]int

func (h *Heap1851) Len() int {
	return len(*h)
}

func (h *Heap1851) Less(i, j int) bool {
	if (*h)[i][0] < (*h)[j][0] {
		return true
	} else if (*h)[i][0] > (*h)[j][0] {
		return false
	} else {
		return (*h)[i][1] <= (*h)[j][1]
	}
}

func (h *Heap1851) Swap(i, j int) {
	(*h)[i], (*h)[j] = (*h)[j], (*h)[i]
}

func (h *Heap1851) Push(x any) {
	*h = append(*h, x.([]int))
}

func (h *Heap1851) Pop() any {
	x := (*h)[len(*h)-1]
	*h = (*h)[:len(*h)-1]
	return x
}

func minInterval(intervals [][]int, queries []int) []int {
	sort.Slice(intervals, func(i, j int) bool {
		if intervals[i][0] < intervals[j][0] {
			return true
		} else if intervals[i][0] > intervals[j][0] {
			return false
		} else {
			return intervals[i][1] <= intervals[j][1]
		}
	})
	queriesWithRank := make([][2]int, len(queries))
	for i, x := range queries {
		queriesWithRank[i] = [2]int{x, i}
	}
	sort.Slice(queriesWithRank, func(i, j int) bool {
		return queriesWithRank[i][0] < queriesWithRank[j][0]
	})
	h := &Heap1851{}
	heap.Init(h)
	res := make([]int, len(queries))
	i := 0
	for _, x := range queriesWithRank {
		for i < len(intervals) && intervals[i][0] <= x[0] {
			heap.Push(h, []int{intervals[i][1] - intervals[i][0] + 1, intervals[i][1]})
			i++
		}
		for len(*h) > 0 {
			v := (*h)[0]
			if v[1] < x[0] {
				heap.Pop(h)
			} else {
				break
			}
		}
		if len(*h) > 0 {
			v := (*h)[0]
			res[x[1]] = v[0]
		} else {
			res[x[1]] = -1
		}
	}
	return res
}

func main() {
	intervals := [][]int{{2, 3}, {2, 5}, {1, 8}, {20, 25}}
	queries := []int{2, 19, 5, 22}
	fmt.Println(minInterval(intervals, queries))
}
