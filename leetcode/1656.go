/**
 * @author: baowj
 * @time: 2022/8/16 10:11:33
 */
package main

type StreamNode struct {
	key   int
	value string
	next  *StreamNode
}

type OrderedStream struct {
	length    int
	maxLength int
	ptr       int
	head      *StreamNode
}

func Constructor1656(n int) OrderedStream {
	return OrderedStream{
		length:    0,
		maxLength: n,
		head:      nil,
		ptr:       1,
	}
}

func (this *OrderedStream) Insert(idKey int, value string) []string {
	if this.length == this.maxLength {
		return this.Pop()
	}
	node := StreamNode{
		key:   idKey,
		value: value,
		next:  nil,
	}
	if this.head == nil {
		this.head = &node
	} else if this.head.key > idKey {
		node.next = this.head
		this.head = &node
	} else {
		tmp := this.head
		for tmp.next != nil && tmp.next.key < idKey {
			tmp = tmp.next
		}
		if tmp.next == nil {
			tmp.next = &node
		} else {
			node.next = tmp.next
			tmp.next = &node
		}
	}
	this.length++
	return this.Pop()
}

func (this *OrderedStream) Pop() []string {
	res := []string{}
	for this.head != nil && this.head.key == this.ptr {
		res = append(res, this.head.value)
		this.head = this.head.next
		this.ptr++
		this.length--
	}
	return res
}
