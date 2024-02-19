package main

func stoneGameVII(stones []int) int {
	dp := make([][]int, len(stones))
	sum := make([]int, len(stones))
	sum[0] = stones[0]
	for i := 0; i < len(stones); i++ {
		dp[i] = make([]int, len(stones))
		for j := 0; j < len(stones); j++ {
			dp[i][j] = -1
		}
		dp[i][i] = 0
		if i > 0 {
			sum[i] = sum[i-1] + stones[i]
		}
	}

	var f func(i, j int) int
	f = func(i, j int) int {
		if dp[i][j] != -1 {
			return dp[i][j]
		}
		l := sum[j] - sum[i] - f(i+1, j)
		r := sum[j-1] - f(i, j-1)
		if i > 0 {
			r -= sum[i-1]
		}
		if r < l {
			dp[i][j] = l
		} else {
			dp[i][j] = r
		}
		return dp[i][j]

	}

	res := f(0, len(stones)-1)
	return res
}
