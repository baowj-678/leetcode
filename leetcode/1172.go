/**
 * @author: baowj
 * @time: 2023/4/28 9:50
 */
package main

import (
	"container/list"
	"fmt"
)

type DinnerPlates struct {
	plates []*[]int
	cap    int
	head   *list.List
	//m      map[int]*list.Element
}

func Constructor(capacity int) DinnerPlates {
	list := list.New()
	list.PushBack(0)
	return DinnerPlates{
		plates: make([]*[]int, 100000),
		head:   list,
		cap:    capacity,
	}
}

func (this *DinnerPlates) Push(val int) {
	head := this.head
	if head.Len() == 0 {
		panic("error")
	}
	element := head.Front()
	idx := element.Value.(int)

	isfull := this.pushPlate(idx, val)
	if isfull {
		// remove list idx
		//delete(this.m, idx)
		head.Remove(head.Front())
		if head.Len() == 0 {
			head.PushBack(idx + 1)
			//this.m[idx+1] = head.Back()
		}
	}
}

func (this *DinnerPlates) pushPlate(idx int, val int) (isfull bool) {
	plate := this.plates[idx]
	if plate == nil {
		plate = &[]int{}
	}
	*plate = append(*plate, val)
	if len(*plate) >= this.cap {
		isfull = true
	}
	this.plates[idx] = plate
	return
}

func (this *DinnerPlates) Pop() int {
	list := this.head
	tail := list.Back().Value.(int)
	list.Remove(list.Back())
	for ; tail >= 0 && (this.plates[tail] == nil || len(*this.plates[tail]) == 0); tail-- {
		if list.Back() != nil && list.Back().Value.(int) >= tail {
			list.Remove(list.Back())
		}
	}
	if tail == -1 {
		list.PushBack(0)
		return -1
	}
	plate := this.plates[tail]
	res := (*plate)[len(*plate)-1]
	list.PushBack(tail)
	*plate = (*plate)[:len(*plate)-1]
	return res
}

func (this *DinnerPlates) PopAtStack(index int) int {
	plate := this.plates[index]
	if plate == nil {
		return -1
	}
	if len(*plate) == 0 {
		return -1
	}
	res := (*plate)[len(*plate)-1]
	if len(*plate) == this.cap {
		insertIdx(this.head, index)
	}
	*plate = (*plate)[:len(*plate)-1]
	return res
}

func insertIdx(list *list.List, idx int) {
	head := list.Front()
	for ; head != nil; head = head.Next() {
		if head.Value.(int) > idx {
			list.InsertBefore(idx, head)
			break
		}
	}
	if head == nil {
		list.PushBack(idx)
	}
}

func main() {
	D := Constructor(2)
	D.Push(1)
	D.Push(2)
	D.Push(3)
	D.Push(4)
	D.Push(5)

	fmt.Println(D.PopAtStack(0))
	D.Push(20)
	D.Push(21)
	fmt.Println(D.PopAtStack(1))
	fmt.Println(D.PopAtStack(1))

	fmt.Println(D.Pop())
	fmt.Println(D.Pop())
	fmt.Println(D.Pop())
	fmt.Println(D.Pop())
	fmt.Println(D.Pop())
}
