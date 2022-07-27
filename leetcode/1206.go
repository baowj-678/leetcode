/**
 * @author: baowj
 * @time: 2022/7/26 09:51:28
 */
package main

import (
	"math/rand"
)

const MaxLevel int = 16
const P float32 = 0.5

type SkipNode struct {
	value   int
	forward []*SkipNode
}

type Skiplist struct {
	head  *SkipNode
	len   int
	level int
}

// New Create new Skiplist
func Constructor() Skiplist {
	return Skiplist{
		head: &SkipNode{
			-1,
			make([]*SkipNode, MaxLevel),
		},
		level: 0,
		len:   0,
	}
}

// randomLevel get level
func randomLevel() int {
	level := 1
	for rand.Float32() < P && level < MaxLevel {
		level++
	}
	return level
}

func (this *Skiplist) Add(num int) {
	update := make([]*SkipNode, MaxLevel)
	tmpNode := this.head
	for l := this.level - 1; l >= 0; l-- {
		for tmpNode.forward[l] != nil {
			if num > tmpNode.forward[l].value {
				tmpNode = tmpNode.forward[l]
			} else {
				break
			}
		}
		update[l] = tmpNode
	}
	// get random level.
	level := randomLevel()
	// make sure the new level is one bigger than the old level at most.
	if level > this.level {
		level = this.level + 1
		update[this.level] = this.head
		this.level = level
	}
	// create new SkipNode
	newNode := &SkipNode{value: num, forward: make([]*SkipNode, level)}
	for i := 0; i < level; i++ {
		newNode.forward[i] = update[i].forward[i]
		update[i].forward[i] = newNode
	}
}

func (this *Skiplist) Search(target int) bool {
	tmpNode := this.head
	for l := this.level - 1; l >= 0; l-- {
		for tmpNode.forward[l] != nil {
			if target > tmpNode.forward[l].value {
				tmpNode = tmpNode.forward[l]
			} else if target == tmpNode.forward[l].value {
				return true
			} else {
				break
			}
		}
	}
	return false
}

func (this *Skiplist) Erase(num int) bool {
	tmpNode := this.head
	res := false
	for l := this.level - 1; l >= 0; l-- {
		for tmpNode.forward[l] != nil {
			if num > tmpNode.forward[l].value {
				tmpNode = tmpNode.forward[l]
			} else if num == tmpNode.forward[l].value {
				tmpNode.forward[l] = tmpNode.forward[l].forward[l]
				res = true
				break
			} else {
				break
			}
		}
	}
	return res
}

func main() {
}
