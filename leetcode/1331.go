/**
 * @author: baowj
 * @time: 2022/7/28 09:54:03
 */
package main

import (
	"math"
	"sort"
)

type Node struct {
	num int
	ref int
}

type ByNum []Node

func (n ByNum) Len() int {
	return len(n)
}

func (n ByNum) Less(i int, j int) bool {
	return n[i].num < n[j].num
}

func (n ByNum) Swap(i int, j int) {
	n[i], n[j] = n[j], n[i]
}

func arrayRankTransform(arr []int) []int {
	nodeArr := make([]Node, len(arr))
	for i, num := range arr {
		nodeArr[i].ref = i
		nodeArr[i].num = num
	}
	sort.Sort(ByNum(nodeArr))
	i := 0
	prevNum := math.MinInt
	for _, node := range nodeArr {
		if prevNum != node.num {
			i += 1
			prevNum = node.num
		}
		arr[node.ref] = i
	}
	return arr
}

func main() {
	a := []int{2, 3, 6}
	for i, num := range a {
		println(i, "-", num)
	}
}
