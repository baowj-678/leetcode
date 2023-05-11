/**
 * @author: baowj
 * @time: 2023/4/16 17:30:54
 * @email: bwj_678@qq.com
 */
package main

import "sort"

type MajorityChecker struct {
	keyToPositon map[int][]int
	arr          []int
}

func Constructor1157(arr []int) MajorityChecker {
	m := map[int][]int{}
	for i, v := range arr {
		if _, ok := m[v]; !ok {
			m[v] = []int{}
		}
		m[v] = append(m[v], i)
	}
	return MajorityChecker{arr: arr, keyToPositon: m}
}

func (this *MajorityChecker) Query(left int, right int, threshold int) int {
	l := right - left + 1
	flags := make([]bool, l)

	next := func(i int) int {
		for flags[i] == true {
			i = (i + 1) % l
		}
		return i
	}
	length := 30
	if l < 30 {
		length = l
	}

	check := func(v int) bool {
		tmp := this.keyToPositon[v]
		a := sort.Search(len(tmp), func(i int) bool {
			return this.arr[i] >= left
		})
		b := sort.Search(len(tmp), func(i int) bool {
			return this.arr[i] > right
		})
		return (b - a) >= threshold
	}
	for i := 0; i < length; i++ {
		idx := next(i)
		if check(this.arr[left+idx]) {
			return this.arr[left+idx]
		}
		flags[idx] = true
	}
	return -1
}

/**
 * Your MajorityChecker object will be instantiated and called as such:
 * obj := Constructor(arr);
 * param_1 := obj.Query(left,right,threshold);
 */
