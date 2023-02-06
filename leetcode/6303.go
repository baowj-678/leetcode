/**
 * @author: baowj
 * @time: 2023/2/5 11:50:03
 * @email: bwj_678@qq.com
 */
package main

func separateDigits(nums []int) []int {
	res := []int{}

	subfunc := func(num int) []int {
		ans := []int{}
		for num > 0 {
			ans = append(ans, num%10)
			num /= 10
		}
		l, r := 0, len(ans)-1
		for l < r {
			ans[l], ans[r] = ans[r], ans[l]
			l++
			r--
		}
		return ans
	}

	for _, num := range nums {
		res = append(res, subfunc(num)...)
	}
	return res
}
