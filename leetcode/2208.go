/**
 * @author: baowj
 * @time: 2023/7/25 10:02
 */
package main

import "container/heap"

type PriorityQueue2208 []float64

func (p *PriorityQueue2208) Len() int {
	return len(*p)
}

func (p *PriorityQueue2208) Less(i, j int) bool {
	return (*p)[i] > (*p)[j]
}

func (p *PriorityQueue2208) Swap(i, j int) {
	(*p)[i], (*p)[j] = (*p)[j], (*p)[i]
}

func (p *PriorityQueue2208) Push(x any) {
	*p = append(*p, x.(float64))
}

func (p *PriorityQueue2208) Pop() any {
	res := (*p)[len(*p)-1]
	*p = (*p)[:len(*p)-1]
	return res
}

func halveArray(nums []int) int {
	h := PriorityQueue2208(make([]float64, len(nums)))
	sum := float64(0)
	minusSum := float64(0)
	for i, x := range nums {
		h[i] = float64(x)
		sum += float64(x)
	}
	heap.Init(&h)
	res := 0
	for minusSum*2 < sum {
		x := heap.Pop(&h).(float64)
		minusSum += x / 2
		heap.Push(&h, x/2)
		res++
	}
	return res
}
