/**
 * @author: baowj
 * @time: 2023/5/21 15:57:49
 * @email: bwj_678@qq.com
 */
package main

import (
	"container/heap"
	"math"
)

type Elem struct {
	b int
	v int
}

func (e *Elem) getTime() int {
	return (e.v + e.b - 1) / e.b
}

type MaxPQ []*Elem

func (q MaxPQ) Len() int {
	return len(q)
}

func (q MaxPQ) Less(i, j int) bool {
	return q[i].getTime() > q[j].getTime()
}

func (q MaxPQ) Swap(i, j int) {
	q[i], q[j] = q[j], q[i]
}

func (q *MaxPQ) Push(x any) {
	e := x.(*Elem)
	*q = append(*q, e)
}

func (q *MaxPQ) Pop() any {
	e := (*q)[len(*q)-1]
	*q = (*q)[:len(*q)-1]
	return e
}

func storeWater(bucket []int, vat []int) int {
	res := 0
	elements := make([]*Elem, len(bucket))
	q := MaxPQ(elements)
	for i := 0; i < len(bucket); i++ {
		e := &Elem{b: bucket[i], v: vat[i]}
		elements[i] = e
		if e.b == 0 && e.v != 0 {
			res++
			e.b = 1
		}
		if e.b == 0 && e.v == 0 {
			e.b = 1
			e.v = 0
		}
	}
	min := math.MaxInt
	lastTime := math.MaxInt
	heap.Init(&q)
	for len(q) > 0 {
		e := heap.Pop(&q).(*Elem)
		t := e.getTime()
		if t < lastTime {
			tmp := t + res
			if tmp >= min {
				return min
			} else {
				min = tmp
			}
		} else {
			res++
			e.b++
		}
		if t <= 1 {
			break
		}
		heap.Push(&q, e)
		lastTime = t
	}
	return min
}

func main() {
	storeWater([]int{0}, []int{0})
}
