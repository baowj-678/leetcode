/**
 * @author: baowj
 * @time: 2022/11/18 10:28
 */
package main

import (
	"sort"
)

var MOD891 int = 1e9 + 7

func sumSubseqWidths(nums []int) int {
	sort.Ints(nums)
	pow := make([]int, len(nums)+1)
	pow[0] = 1
	for i := 1; i < len(pow); i++ {
		pow[i] = (pow[i-1] * 2) % MOD891
	}
	res := 0
	for i := 0; i < len(nums); i++ {
		res = (res + (pow[i]-1)*nums[i]) % MOD891
		res = (res - (pow[len(nums)-i-1]-1)*nums[i]) % MOD891
	}
	return res
}

func main() {
	nums := []int{2, 1, 3}
	print(sumSubseqWidths(nums))
}
