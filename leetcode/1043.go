/**
 * @author: baowj
 * @time: 2023/4/19 12:08:21
 * @email: bwj_678@qq.com
 */
package main

func maxSumAfterPartitioning(arr []int, k int) int {
	dp := make([]int, len(arr))
	for i := 0; i < len(arr); i++ {
		max := 0
		maxA := 0
		for t := 0; t < k && i-t >= 0; t++ {
			if maxA < arr[i-t] {
				maxA = arr[i-t]
			}
			var v int
			if i-t <= 0 {
				v = (t + 1) * maxA
			} else {
				v = (t+1)*maxA + dp[i-t-1]
			}
			if v > max {
				max = v
			}
		}
		dp[i] = max
	}
	return dp[len(dp)-1]
}
