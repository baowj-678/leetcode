/**
 * @author: baowj
 * @time: 2022/4/2 16:23:23
 */

package main

func minDistance(word1 string, word2 string) int {
	dp := [505][505]int{}
	dp[0][0] = 0
	for i := 1; i <= len(word1); i++ {
		dp[i][0] = i
	}
	for i := 1; i <= len(word2); i++ {
		dp[0][i] = i
	}
	for i := 1; i <= len(word1); i++ {
		for j := 1; j <= len(word2); j++ {
			a := dp[i-1][j] + 1
			b := dp[i][j-1] + 1
			c := dp[i-1][j-1]
			if word1[i-1] != word2[j-1] {
				c++
			}
			if a > b {
				a = b
			}
			if a > c {
				a = c
			}
			dp[i][j] = a
		}
	}
	return dp[len(word1)][len(word2)]
}

func main() {
	word1 := "horse"
	word2 := "rostyfftfy"
	print(minDistance(word1, word2))
}
