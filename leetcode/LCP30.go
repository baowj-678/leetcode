package main

import "container/heap"

type HeapInt30 []int

func (h *HeapInt30) Len() int {
	return len(*h)
}

func (h *HeapInt30) Less(i, j int) bool {
	return (*h)[i] < (*h)[j]
}

func (h *HeapInt30) Swap(i, j int) {
	(*h)[i], (*h)[j] = (*h)[j], (*h)[i]
}

func (h *HeapInt30) Push(x any) {
	*h = append(*h, x.(int))
}

func (h *HeapInt30) Pop() any {
	v := (*h)[len(*h)-1]
	*h = (*h)[:len(*h)-1]
	return v
}

func magicTower(nums []int) int {
	heapInt := HeapInt30([]int{})
	tailInt := []int{}
	heap.Init(&heapInt)
	sum := 0
	res := 0
	for _, num := range nums {
		if num < 0 {
			heap.Push(&heapInt, num)
		}
		sum += num

		for sum < 0 {
			if len(heapInt) > 0 {
				head := heap.Pop(&heapInt).(int)
				sum -= head
				tailInt = append(tailInt, head)
				res++
			} else {
				return -1
			}
		}
	}

	for _, num := range tailInt {
		sum += num
		if sum < 0 {
			return -1
		}
	}
	return res
}
