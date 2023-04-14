/**
 * @author: baowj
 * @time: 2023/3/29 12:44
 */
package main

func countVowelStrings(n int) int {
	dp := [5]int{1, 1, 1, 1, 1}
	for i := 1; i < n; i++ {
		for k := 0; k < 5; k++ {
			for t := k + 1; t < 5; t++ {
				dp[k] += dp[t]
			}
		}

	}
	res := 0
	for i := 0; i < len(dp); i++ {
		res += dp[i]
	}
	return res
}

func main() {
	println(countVowelStrings(2))
}
