/**
 * @author: baowj
 * @time: 2022/4/4 11:27:53
 */
package main

import (
	"math"
)

const MaxLen int = 10

type NumArray struct {
	array  []int
	vals   [MaxLen]int
	indexs [MaxLen]int
	len    int
}

func constructor(nums []int) NumArray {
	numArray := new(NumArray)
	for i := 1; i < len(nums); i++ {
		nums[i] += nums[i-1]
	}
	numArray.array = nums
	numArray.vals = [MaxLen]int{}
	numArray.indexs = [MaxLen]int{}
	numArray.len = 0
	for i := 0; i < MaxLen; i++ {
		numArray.vals[i] = math.MaxInt
		numArray.indexs[i] = math.MaxInt
	}
	return *numArray
}

func (this *NumArray) Update(index int, val int) {
	if this.len >= MaxLen {
		this.refresh()
	}
	if index == 0 {
		val = val - this.array[index]
	} else {
		val = val - (this.array[index] - this.array[index-1])
	}
	for i := 0; i < this.len; i++ {
		if index == this.indexs[i] {
			this.vals[i] = val
			index = math.MaxInt
			val = math.MaxInt
			break
		} else if index < this.indexs[i] {
			swap(&index, &(this.indexs[i]))
			swap(&val, &(this.vals[i]))
		}
	}
	if index != math.MaxInt {
		this.indexs[this.len] = index
		this.vals[this.len] = val
		this.len++
	}
}

func swap(a *int, b *int) {
	var tmp int = *a
	*a = *b
	*b = tmp
}

func (this *NumArray) SumRange(left int, right int) int {
	l := math.MaxInt
	for t := 0; t < this.len; t++ {
		if this.indexs[t] >= left {
			l = this.indexs[t]
			break
		}
	}
	r := math.MinInt
	for t := this.len - 1; t >= 0; t-- {
		if this.indexs[t] <= right {
			r = this.indexs[t]
			break
		}
	}
	if l <= r {
		this.refresh()
	}

	if left == 0 {
		return this.array[right]
	} else {
		return this.array[right] - this.array[left-1]
	}
}

func (this *NumArray) refresh() {
	add := 0
	for t := 0; t < this.len; t++ {
		i := this.indexs[t]
		if i == 0 {
			add += this.vals[t]
		} else {
			add += this.vals[t]
		}
		var index int
		if t == this.len-1 {
			index = len(this.array)
		} else {
			index = this.indexs[t+1]
		}
		for ; i < index; i++ {
			this.array[i] += add
		}
		this.indexs[t] = math.MaxInt
		this.vals[t] = math.MaxInt
	}
	this.len = 0
}

/**
 * Your NumArray object will be instantiated and called as such:
 * obj := Constructor(nums);
 * obj.Update(index,val);
 * param_2 := obj.SumRange(left,right);
 */
