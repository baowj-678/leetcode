/**
 * @author: baowj
 * @time: 2023/2/15 18:40
 */
package main

func isGoodArray(nums []int) bool {
	gcd := func(x, y int) int {
		if x < y {
			x, y = y, x
		}
		for y != 0 {
			x, y = y, x/y
		}
		return x
	}
	res := nums[0]
	for _, num := range nums {
		res = gcd(res, num)
	}
	return res == 1
}
