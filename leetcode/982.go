/**
 * @author: baowj
 * @time: 2023/3/5 10:50:58
 * @email: bwj_678@qq.com
 */
package main

func countTriplets(nums []int) int {
	dp := make([][]int, 3)
	for i := 0; i < 3; i++ {
		dp[i] = make([]int, 1<<16)
	}
	for _, num := range nums {
		dp[0][num] += 1
	}

	for i := 1; i <= 2; i++ {
		for k := 0; k < 1<<16; k++ {
			for _, num := range nums {
				dp[i][k&num] += dp[i-1][k]
			}
		}
	}
	return dp[2][0]
}
