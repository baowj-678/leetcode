/**
 * @author: baowj
 * @time: 2022/4/13 10:07:48
 */
package main

import (
	"fmt"
	"math/rand"
)

type RandomizedSet struct {
	Map   map[int]int
	array []int
}

func Constructor() RandomizedSet {
	return RandomizedSet{
		map[int]int{},
		[]int{},
	}
}

func (this *RandomizedSet) Insert(val int) bool {
	_, ok := this.Map[val]
	if ok {
		return false
	} else {
		this.Map[val] = len(this.array)
		this.array = append(this.array, val)
		return true
	}
}

func (this *RandomizedSet) Remove(val int) bool {
	v, ok := this.Map[val]
	if ok {
		delete(this.Map, val)
		if v != len(this.array)-1 {
			this.array[v] = this.array[len(this.array)-1]
			this.Map[this.array[v]] = v
		}
		this.array = this.array[:len(this.array)-1]
		return true
	} else {
		return false
	}
}

func (this *RandomizedSet) GetRandom() int {
	intn := rand.Intn(len(this.array))
	return this.array[intn]
}

func main() {
	obj := Constructor()
	obj.Insert(0)
	obj.Remove(0)
	fmt.Println(obj.Insert(0))
}

/**
 * Your RandomizedSet object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.Insert(val);
 * param_2 := obj.Remove(val);
 * param_3 := obj.GetRandom();
 */
