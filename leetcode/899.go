/**
 * @author: baowj
 * @time: 2022/8/3 09:40:08
 */
package main

import "sort"

type IncreaseOrder []uint8

func (this IncreaseOrder) Len() int {
	return len(this)
}

func (this IncreaseOrder) Less(i int, j int) bool {
	return this[i] < this[j]
}

func (this IncreaseOrder) Swap(i int, j int) {
	this[i], this[j] = this[j], this[i]
}

func orderlyQueue(s string, k int) string {
	var res string
	if k == 1 {
		maxStart := 0
		for i := 1; i < len(s); i++ {
			if cmp899(s, maxStart, i) > 0 {
				maxStart = i
			}
		}
		res = s[maxStart:] + s[:maxStart]
	} else {
		newStr := make([]uint8, len(s))
		for i := 0; i < len(s); i++ {
			newStr[i] = s[i]
		}
		sort.Sort(IncreaseOrder(newStr))
		res = string(newStr)
	}
	return res
}

func cmp899(s string, startA int, startB int) int {
	for i := 0; i < len(s); i++ {
		a := (startA + i) % len(s)
		b := (startB + i) % len(s)
		if s[a] < s[b] {
			return -1
		} else if s[a] > s[b] {
			return 1
		}
	}
	return 0
}

func main() {
	s := "cba"
	println(orderlyQueue(s, 1))
}
