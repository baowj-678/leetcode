/**
 * @author: baowj
 * @time: 2022/8/26 09:56:19
 */
package main

func maxProduct(nums []int) int {
	ans := [2]int{0, 0}
	for _, x := range nums {
		if x >= ans[1] {
			ans[0], ans[1] = ans[1], x
		} else if x >= ans[0] {
			ans[0] = x
		}
	}
	return (ans[0] - 1) * (ans[1] - 1)
}
