/**
 * @author: baowj
 * @time: 2023/2/8 21:51:23
 * @email: bwj_678@qq.com
 */
package main

func findIntegers(n int) int {
	table := [33]int{}
	c := func(n, k int) int {
		a, b := 1, n
		var res int64 = 1
		for {
			if a > k {
				res *= int64(b)
				b--
			} else {
				if res%int64(a) == 0 {
					res /= int64(a)
					a++
				} else {
					res *= int64(b)
					b--
				}
			}
			if a > k && b < n-k+1 {
				break
			}
		}
		return int(res)
	}

	for i := 1; i < 33; i++ {
		v := 1
		for k := 1; k <= (i+1)/2; k++ {
			v += c(i+1-k, k)
		}
		table[i] = v
	}

	nums := []int{}
	for n > 0 {
		if n&0b1 == 0b1 {
			nums = append(nums, 1)
		} else {
			nums = append(nums, 0)
		}
		n >>= 1
	}
	dp := make([]int, len(nums))
	if nums[0] == 1 {
		dp[0] = 2
	} else {
		dp[0] = 1
	}
	for i := 1; i < len(dp); i++ {
		if nums[i] == 1 {
			dp[i] = table[i]
			if nums[i-1] == 0 {
				dp[i] += dp[i-1]
			} else {
				dp[i] += table[i-1]
			}
		} else {
			dp[i] = dp[i-1]
		}
	}
	return dp[len(dp)-1]
}
