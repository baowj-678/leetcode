package main

func minExtraChar(s string, dictionary []string) int {
	dp := make([]int, len(s))
	for i := 0; i < len(dp); i++ {
		dp[i] = i + 1
	}
	minF := func(x, y int) int {
		if x > y {
			x = y
		}
		return x
	}

	for i := 0; i < len(dp); i++ {
		min := dp[i]
		for j := 0; j < i; j++ {
			min = minF(min, dp[j]+i-j)
		}
		for _, word := range dictionary {
			if len(word) <= i+1 && s[i-len(word)+1:i+1] == word {
				if len(word) == i+1 {
					min = 0
				} else {
					min = minF(min, dp[i-len(word)])
				}
			}
		}
		dp[i] = min
	}
	return dp[len(dp)-1]
}
