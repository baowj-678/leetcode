/**
 * @author: baowj
 * @time: 2023/7/24 10:06
 */
package main

import (
	"container/heap"
	"fmt"
	"math"
)

type PriorityQueue1499 [][]int

func (p *PriorityQueue1499) Len() int {
	return len(*p)
}

func (p *PriorityQueue1499) Less(i, j int) bool {
	return (*p)[i][0] < (*p)[j][0]
}

func (p *PriorityQueue1499) Swap(i, j int) {
	(*p)[i], (*p)[j] = (*p)[j], (*p)[i]
}

func (p *PriorityQueue1499) Push(x any) {
	*p = append(*p, x.([]int))
}

func (p *PriorityQueue1499) Pop() any {
	res := (*p)[p.Len()-1]
	*p = (*p)[:p.Len()-1]
	return res
}

func findMaxValueOfEquation(points [][]int, k int) int {
	h := PriorityQueue1499{}
	heap.Init(&h)

	res := math.MinInt

	for _, point := range points {
		tmp := point[0] + point[1]
		for h.Len() > 0 && point[0]-h[0][1] > k {
			heap.Pop(&h)
		}
		if h.Len() > 0 {
			tmp -= h[0][0]
			if tmp > res {
				res = tmp
			}
		}
		// push
		heap.Push(&h, []int{point[0] - point[1], point[0]})
	}
	return res
}

func main() {
	points := [][]int{{1, 3}, {2, 0}, {5, 10}, {6, -10}}
	k := 1
	fmt.Println(findMaxValueOfEquation(points, k))
}
