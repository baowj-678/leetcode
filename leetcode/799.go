/**
 * @author: baowj
 * @time: 2022/11/21 11:01
 */
package main

import "fmt"

func champagneTower(poured int, query_row int, query_glass int) float64 {
	dp := make([]float64, query_row+2)
	dp[0] = float64(poured)
	for i := 1; i <= query_row; i++ {
		tmp := make([]float64, query_row+2)
		for j := 0; j <= i; j++ {
			if j > 0 && dp[j-1] > 1 {
				tmp[j] += (dp[j-1] - 1) / 2
			}
			if j < i && dp[j] > 1 {
				tmp[j] += (dp[j] - 1) / 2
			}
		}
		dp = tmp
	}
	if dp[query_glass] > 1 {
		return 1
	}
	return dp[query_glass]
}

func main() {
	fmt.Println(champagneTower(3, 2, 2))
}
