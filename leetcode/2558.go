/**
 * @author: baowj
 * @time: 2023/2/6 21:46:05
 * @email: bwj_678@qq.com
 */
package main

import (
	"container/heap"
	"math"
)

type Ints2558 []int

func (s Ints2558) Len() int {
	return len(s)
}

func (s Ints2558) Less(i, j int) bool {
	return s[i] > s[j]
}

func (s Ints2558) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

func (s *Ints2558) Push(x interface{}) {
	*s = append(*s, x.(int))
}

func (s *Ints2558) Pop() interface{} {
	item := (*s)[len(*s)-1]
	*s = (*s)[:len(*s)-1]
	return item
}

func pickGifts(gifts []int, k int) int64 {
	h := (*Ints2558)(&gifts)
	heap.Init(h)
	var res int64 = 0
	for i := 0; i < k; i++ {
		item := heap.Pop(h).(int)
		item = int(math.Floor(math.Sqrt((float64(item)))))
		if item != 0 {
			heap.Push(h, item)
		}
	}
	for _, item := range *(*[]int)(h) {
		res += int64(item)
	}
	return res
}
