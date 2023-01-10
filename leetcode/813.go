/**
 * @author: baowj
 * @time: 2022/12/3 11:27:10
 * @email: bwj_678@qq.com
 */
package main

func largestSumOfAverages(nums []int, k int) float64 {
	dp := make([][]float64, len(nums))
	presum := make([]int, len(nums))
	presum[0] = nums[0]
	for i := 1; i < len(nums); i++ {
		presum[i] = presum[i-1] + nums[i]
	}
	for i := 0; i < len(nums); i++ {
		dp[i] = make([]float64, k)
		dp[i][0] = float64(presum[i]) / float64(i+1)
		for j := 1; j < k && j <= i; j++ {
			max := .0
			for t := j - 1; t < i; t++ {
				println(t, j)
				tmp := dp[t][j-1] + float64(presum[i]-presum[t])/float64(i-t)
				if tmp > max {
					max = tmp
				}
			}
			dp[i][j] = max
		}
	}
	return dp[len(nums)-1][k-1]
}
