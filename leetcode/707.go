/**
 * @author: baowj
 * @time: 2022/9/23 09:31:52
 */
package main

type LinkedListNode struct {
	val  int
	next *LinkedListNode
	prev *LinkedListNode
}

type MyLinkedList struct {
	head   *LinkedListNode
	tail   *LinkedListNode
	length int
}

func Constructor707() MyLinkedList {
	return MyLinkedList{
		head:   nil,
		tail:   nil,
		length: 0,
	}
}

func (this *MyLinkedList) Get(index int) int {
	if index >= this.length {
		return -1
	}
	tmp := this.head
	for tmp != nil && index > 0 {
		tmp = tmp.next
		index--
	}
	return tmp.val
}

func (this *MyLinkedList) AddAtHead(val int) {
	node := &LinkedListNode{
		val:  val,
		next: nil,
		prev: nil,
	}
	if this.length == 0 {
		this.head = node
		this.tail = node
	} else {
		node.next = this.head
		this.head.prev = node
		this.head = node
	}
	this.length++
}

func (this *MyLinkedList) AddAtTail(val int) {
	node := &LinkedListNode{
		val:  val,
		next: nil,
		prev: nil,
	}
	if this.length == 0 {
		this.head = node
		this.tail = node
	} else {
		this.tail.next = node
		node.prev = this.tail
		this.tail = node
	}
	this.length++
}

func (this *MyLinkedList) AddAtIndex(index int, val int) {
	if index <= 0 {
		this.AddAtHead(val)
	} else if index == this.length {
		this.AddAtTail(val)
	} else if index < this.length {
		node := this.head
		for node != nil && index > 0 {
			node = node.next
			index--
		}
		newNode := &LinkedListNode{
			val:  val,
			next: nil,
			prev: nil,
		}
		newNode.next, newNode.prev, node.prev, node.prev.next = node, node.prev, newNode, newNode
		this.length++
	}
}

func (this *MyLinkedList) DeleteAtIndex(index int) {
	if this.length == 0 {
		return
	}
	if index < 0 || index >= this.length {
		return
	}

	if index == 0 || index == this.length-1 {
		if index == 0 {
			this.head = this.head.next
			if this.head != nil {
				this.head.prev = nil
			}
		}
		if index == this.length-1 {
			this.tail = this.tail.prev
			if this.tail != nil {
				this.tail.next = nil
			}
		}
	} else {
		node := this.head
		for node != nil && index > 0 {
			node = node.next
			index--
		}
		node.prev.next, node.next.prev = node.next, node.prev
	}
	this.length--
}

/**
 * Your MyLinkedList object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.Get(index);
 * obj.AddAtHead(val);
 * obj.AddAtTail(val);
 * obj.AddAtIndex(index,val);
 * obj.DeleteAtIndex(index);
 */
