/**
 * @author: baowj
 * @time: 2022/7/16 09:35:37
 */
package main

type MovingAverage struct {
	sum   int
	array []int
	size  int
}

/** Initialize your data structure here. */
func Constructor041(size int) MovingAverage {
	ma := MovingAverage{
		sum:   0,
		array: []int{},
		size:  size,
	}
	return ma
}

func (this *MovingAverage) Next(val int) float64 {
	this.sum = this.sum + val
	if len(this.array) == this.size {
		this.sum -= this.array[0]
		this.array = this.array[1:]
	}
	this.array = append(this.array, val)
	return float64(this.sum) / float64(len(this.array))
}

/**
 * Your MovingAverage object will be instantiated and called as such:
 * obj := Constructor(size);
 * param_1 := obj.Next(val);
 */

func main() {
}
