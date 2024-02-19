package main

import (
	"container/heap"
	"math"
)

type MaxHeap1696 [][]int

func (m *MaxHeap1696) Len() int {
	return len(*m)
}

func (m *MaxHeap1696) Less(i, j int) bool {
	return (*m)[i][0] >= (*m)[j][0]
}

func (m *MaxHeap1696) Swap(i, j int) {
	(*m)[i], (*m)[j] = (*m)[j], (*m)[i]
}

func (m *MaxHeap1696) Push(x any) {
	*m = append(*m, x.([]int))
}

func (m *MaxHeap1696) Pop() any {
	x := (*m)[len(*m)-1]
	*m = (*m)[:len(*m)-1]
	return x
}

func maxResult(nums []int, k int) int {
	dp := make([]int, len(nums))
	dp[0] = nums[0]
	maxHeap := MaxHeap1696([][]int{{dp[0], 0}})
	heap.Init(&maxHeap)

	maxF := func(x, y int) int {
		if x < y {
			x = y
		}
		return x
	}

	for i := 1; i < len(nums); i++ {
		maxValue := math.MinInt
		for len(maxHeap) > 0 {
			if maxHeap[0][1] >= i-k {
				maxValue = maxF(maxValue, maxHeap[0][0])
				break
			} else {
				heap.Pop(&maxHeap)
			}
		}
		dp[i] = nums[i] + maxValue
		heap.Push(&maxHeap, []int{dp[i], i})
	}
	return dp[len(dp)-1]
}
