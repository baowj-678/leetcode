/**
 * @author: baowj
 * @time: 2023/7/11 10:13
 */
package main

func maxAlternatingSum(nums []int) int64 {
	res := int64(0)
	odd, even := int64(0), int64(0)

	max := func(x, y int64) int64 {
		if x < y {
			x = y
		}
		return x
	}

	for _, num := range nums {
		num64 := int64(num)
		odd, even = max(odd, even-num64), max(num64, max(even, odd+num64))
		res = max(res, max(odd, even))
	}
	return res
}
