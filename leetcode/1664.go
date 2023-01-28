/**
 * @author: baowj
 * @time: 2023/1/28 11:23:32
 * @email: bwj_678@qq.com
 */
package main

func waysToMakeFair(nums []int) int {
	even := make([]int, len(nums))
	odd := make([]int, len(nums))
	lastEven, lastOdd := 0, 0
	for i, num := range nums {
		if i%2 == 0 {
			lastEven += num
		} else {
			lastOdd += num
		}
		even[i] = lastEven
		odd[i] = lastOdd
	}
	res := 0
	if even[len(nums)-1]-even[0] == odd[len(nums)-1] {
		res++
	}
	for i := 1; i < len(nums); i++ {
		if even[i-1]+odd[len(nums)-1]-odd[i] == odd[i-1]+even[len(nums)-1]-even[i] {
			res++
		}
	}
	return res
}
