/**
 * @author: baowj
 * @time: 2022/8/2 09:56:17
 */
package main

type MyCircularQueue struct {
	queue   []int
	rear    int
	front   int
	isEmpty bool
}

func Constructor622(k int) MyCircularQueue {
	return MyCircularQueue{
		queue:   make([]int, k),
		rear:    0,
		front:   0,
		isEmpty: true,
	}
}

func (this *MyCircularQueue) EnQueue(value int) bool {
	if this.isEmpty {
		this.isEmpty = false
	} else if this.rear == this.front {
		return false
	}
	this.queue[this.rear] = value
	this.rear = (this.rear + 1) % len(this.queue)
	return true
}

func (this *MyCircularQueue) DeQueue() bool {
	if this.isEmpty {
		return false
	}
	this.front = (this.front + 1) % len(this.queue)
	if this.front == this.rear {
		this.isEmpty = true
	}
	return true
}

func (this *MyCircularQueue) Front() int {
	if this.isEmpty {
		return -1
	}
	return this.queue[this.front]
}

func (this *MyCircularQueue) Rear() int {
	if this.isEmpty {
		return -1
	}
	return this.queue[(this.rear-1+len(this.queue))%len(this.queue)]
}

func (this *MyCircularQueue) IsEmpty() bool {
	return this.isEmpty
}

func (this *MyCircularQueue) IsFull() bool {
	return !this.isEmpty && this.front == this.rear
}
