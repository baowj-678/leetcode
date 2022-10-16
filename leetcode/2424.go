/**
 * @author: baowj
 * @time: 2022/10/11 22:38:30
 */
package main

type LUPrefix struct {
	nums   []int
	prefix int
}

func Constructor(n int) LUPrefix {
	return LUPrefix{
		nums:   make([]int, n+1),
		prefix: 0,
	}
}

func (this *LUPrefix) Upload(video int) {
	this.nums[video] = video
	for ; this.prefix < len(this.nums); this.prefix++ {
		if this.nums[this.prefix] != this.prefix {
			break
		}
	}
	if this.prefix == len(this.nums) {
		this.prefix--
	} else {
		if this.nums[this.prefix] != this.prefix {
			this.prefix--
		}
	}
}

func (this *LUPrefix) Longest() int {
	return this.prefix
}
