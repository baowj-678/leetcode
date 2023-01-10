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

func Constructor() FreqStack {
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

func main() {
	freqStack := Constructor()
	freqStack.Push(5)        //堆栈为 [5]
	freqStack.Push(7)        //堆栈是 [5,7]
	freqStack.Push(5)        //堆栈是 [5,7,5]
	freqStack.Push(7)        //堆栈是 [5,7,5,7]
	freqStack.Push(4)        //堆栈是 [5,7,5,7,4]
	freqStack.Push(5)        //堆栈是 [5,7,5,7,4,5]
	print(freqStack.Pop())   //返回 5 ，因为 5 出现频率最高。堆栈变成 [5,7,5,7,4]。
	print(freqStack.Pop())   //返回 7 ，因为 5 和 7 出现频率最高，但7最接近顶部。堆栈变成 [5,7,5,4]。
	print(freqStack.Pop())   //返回 5 ，因为 5 出现频率最高。堆栈变成 [5,7,4]。
	println(freqStack.Pop()) //返回 4 ，因为 4, 5 和 7 出现频率最高，但 4 是最接近顶部的。堆栈变成 [5,7]
}
