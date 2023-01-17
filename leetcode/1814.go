/**
 * @author: baowj
 * @time: 2023/1/17 20:51:17
 * @email: bwj_678@qq.com
 */
package main

import (
	"sort"
	"strconv"
)

const Mod1814 int = 1e9 + 7

func countNicePairs(nums []int) int {
	for i, num := range nums {
		nums[i] -= rev1814(num)
	}
	sort.Ints(nums)
	last := 0
	res := 0
	for i := 0; i < len(nums); i++ {
		if nums[i] != nums[last] {
			length := i - last
			res = (res + length*(length-1)/2) % Mod1814
			last = i
		}
	}
	length := len(nums) - last
	res = (res + length*(length-1)/2) % Mod1814
	return res
}

func rev1814(num int) int {
	a := []byte(strconv.Itoa(num))
	for i, j := 0, len(a)-1; i < j; {
		a[i], a[j] = a[j], a[i]
		i++
		j--
	}
	x, _ := strconv.Atoi(string(a))
	return x
}

func main() {
	println(rev1814(900))
}
