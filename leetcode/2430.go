/**
 * @author: baowj
 * @time: 2022/10/3 16:51:11
 */
package main

func deleteString(s string) int {
	dp := make([]int, len(s))
	for i := 0; i < len(dp); i++ {
		dp[i] = -1
	}
	for i := 1; i <= len(s)/2; i++ {
		isSame := true
		for l, r := 0, i; l < i; l, r = l+1, r+1 {
			if s[l] != s[r] {
				isSame = false
				break
			}
		}
		if isSame {
			dp[i-1] = 1
		}
	}
	for i := 1; i < len(s); i++ {
		if dp[i-1] != -1 {
			for j := 1; i+j*2 <= len(s); j++ {
				isSame := true
				for l, r := i, i+j; l < i+j; l, r = l+1, r+1 {
					if s[l] != s[r] {
						isSame = false
						break
					}
				}
				if isSame {
					if dp[i+j-1] < dp[i-1]+1 {
						dp[i+j-1] = dp[i-1] + 1
					}
				}
			}
		}
	}
	max := 0
	for i := 0; i < len(dp); i++ {
		if max < dp[i] {
			max = dp[i]
		}
	}
	return max + 1
}
