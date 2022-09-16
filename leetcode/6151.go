/**
 * @author: baowj
 * @time: 2022/8/14 10:52:01
 */
package main

func countSpecialNumbers(n int) int {
	nums := []int{}
	for n > 0 {
		nums = append(nums, n%10)
		n /= 10
	}
	for i, j := 0, len(nums)-1; i < j; {
		nums[i], nums[j] = nums[j], nums[i]
		j--
		i++
	}
	return dp(0, nums, 0, false, true)
}

/*
i: 当前位数
nums: 数字各位组成的数组
mask: 已经包含的数字
isNum: 第i位之前是否填过数字
isLimit: 第i位之前填的是否都是nums中的数字
*/
func dp(i int, nums []int, mask int, isNum bool, isLimit bool) int {
	if i == len(nums) {
		if !isNum {
			return 0
		}
		return 1
	}
	res := 0
	from, to := 0, 9
	if !isNum {
		// i前面都没填
		res += dp(i+1, nums, 0, false, false)
		from = 1
	}
	if isLimit {
		to = nums[i]
	}
	// 该位填什么
	for j := from; j <= to; j++ {
		// 没填过
		if mask&(1<<j) == 0 {
			res += dp(i+1, nums, mask|(1<<j), true, isLimit && (j == to))
		}
	}
	return res
}

func main() {
	print(countSpecialNumbers(20))
}
