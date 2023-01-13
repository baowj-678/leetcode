/**
 * @author: baowj
 * @time: 2022/12/3 11:35:16
 * @email: bwj_678@qq.com
 */
package main

type FreqStack struct {
	maxFreq  int
	num2freq map[int]int
	stacks   [][]int
}

func Constructor895() FreqStack {
	return FreqStack{
		maxFreq:  0,
		num2freq: map[int]int{},
		stacks:   [][]int{},
	}
}

func (this *FreqStack) Push(val int) {
	if freq, ok := this.num2freq[val]; ok {
		freq++
		this.num2freq[val] = freq
		if freq > len(this.stacks) {
			this.stacks = append(this.stacks, []int{val})
		} else {
			this.stacks[freq-1] = append(this.stacks[freq-1], val)
		}
	} else {
		this.num2freq[val] = 1
		if len(this.stacks) == 0 {
			this.stacks = append(this.stacks, []int{})
		}
		this.stacks[0] = append(this.stacks[0], val)
	}
}

func (this *FreqStack) Pop() int {
	length := len(this.stacks) - 1
	stack := this.stacks[length]
	res := stack[len(stack)-1]
	this.num2freq[res]--
	if len(stack) == 1 {
		this.stacks = this.stacks[:length]
	} else {
		this.stacks[length] = stack[:len(stack)-1]
	}
	return res
}

/**
 * Your FreqStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(val);
 * param_2 := obj.Pop();
 */
