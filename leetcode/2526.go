/**
 * @author: baowj
 * @time: 2023/1/13 11:49:04
 * @email: bwj_678@qq.com
 */
package main

type DataStream struct {
	length int
	k      int
	value  int
}

func Constructor(value int, k int) DataStream {
	return DataStream{
		length: 0,
		k:      k,
		value:  value,
	}
}

func (this *DataStream) Consec(num int) bool {
	if num == this.value {
		this.length++
		if this.length >= this.k {
			return true
		} else {
			return false
		}
	}
	this.length = 0
	return false
}

/**
 * Your DataStream object will be instantiated and called as such:
 * obj := Constructor(value, k);
 * param_1 := obj.Consec(num);
 */
