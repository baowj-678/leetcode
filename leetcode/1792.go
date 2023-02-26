/**
 * @author: baowj
 * @time: 2023/2/19 11:58:29
 * @email: bwj_678@qq.com
 */
package main

import h "container/heap"

type MyStruct struct {
	score *[]float32
	data  *[][]int
	heap  *[]int
}

func (this MyStruct) Len() int {
	return len(*this.heap)
}

func (this MyStruct) Less(i, j int) bool {
	return (*this.score)[(*this.heap)[i]] >= (*this.score)[(*this.heap)[j]]
}

func (this MyStruct) Swap(i, j int) {
	(*this.heap)[i], (*this.heap)[j] = (*this.heap)[j], (*this.heap)[i]
}

func (this *MyStruct) Push(x interface{}) {
	idx := x.(int)
	*this.heap = append(*this.heap, idx)
}

func (this *MyStruct) Pop() interface{} {
	x := (*this.heap)[len(*this.heap)-1]
	*this.heap = (*this.heap)[:len(*this.heap)-1]
	return x
}

func getScore(a, n int) float32 {
	return float32(n-a) / (float32(n) * float32(n+1))
}

func maxAverageRatio(classes [][]int, extraStudents int) float64 {
	score := make([]float32, len(classes))
	heap := make([]int, len(classes))
	stru := MyStruct{
		score: &score,
		data:  &classes,
		heap:  &heap,
	}
	for i := 0; i < len(classes); i++ {
		(*stru.heap)[i] = i
		(*stru.score)[i] = getScore((*stru.data)[i][0], (*stru.data)[i][1])
	}
	h.Init(&stru)

	for i := 0; i < extraStudents; i++ {
		x := h.Pop(&stru)
		(*stru.data)[x.(int)][0]++
		(*stru.data)[x.(int)][1]++
		(*stru.score)[x.(int)] = getScore((*stru.data)[x.(int)][0], (*stru.data)[x.(int)][1])
		h.Push(&stru, x)
	}
	res := 0.0
	for _, class := range *stru.data {
		res += float64(class[0]) / float64(class[1])
	}
	res /= float64(len(classes))
	return res
}
