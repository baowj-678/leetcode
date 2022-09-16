/**
 * @author: baowj
 * @time: 2022/8/15 10:12:56
 */
package main

type MyCircularDeque struct {
	queue []int
	k     int
}

func Constructor641(k int) MyCircularDeque {
	return MyCircularDeque{
		queue: []int{},
		k:     k,
	}
}

func (this *MyCircularDeque) InsertFront(value int) bool {
	if len(this.queue) == this.k {
		return false
	}
	this.queue = append([]int{value}, this.queue...)
	return true
}

func (this *MyCircularDeque) InsertLast(value int) bool {
	if len(this.queue) == this.k {
		return false
	}
	this.queue = append(this.queue, value)
	return true
}

func (this *MyCircularDeque) DeleteFront() bool {
	if len(this.queue) == 0 {
		return false
	}
	this.queue = this.queue[1:]
	return true
}

func (this *MyCircularDeque) DeleteLast() bool {
	if len(this.queue) == 0 {
		return false
	}
	this.queue = this.queue[:len(this.queue)-1]
	return true
}

func (this *MyCircularDeque) GetFront() int {
	if len(this.queue) == 0 {
		return -1
	}
	return this.queue[0]
}

func (this *MyCircularDeque) GetRear() int {
	if len(this.queue) == 0 {
		return -1
	}
	return this.queue[len(this.queue)-1]
}

func (this *MyCircularDeque) IsEmpty() bool {
	return len(this.queue) == 0
}

func (this *MyCircularDeque) IsFull() bool {
	return len(this.queue) == this.k
}

/**
 * Your MyCircularDeque object will be instantiated and called as such:
 * obj := Constructor(k);
 * param_1 := obj.InsertFront(value);
 * param_2 := obj.InsertLast(value);
 * param_3 := obj.DeleteFront();
 * param_4 := obj.DeleteLast();
 * param_5 := obj.GetFront();
 * param_6 := obj.GetRear();
 * param_7 := obj.IsEmpty();
 * param_8 := obj.IsFull();
 */
