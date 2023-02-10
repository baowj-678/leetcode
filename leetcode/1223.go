/**
 * @author: baowj
 * @time: 2023/2/10 11:10:20
 * @email: bwj_678@qq.com
 */
package main

import "fmt"

func dieSimulator(n int, rollMax []int) int {
	const mod int64 = 1e9 + 7
	dp := make([][6]int, n)
	dp[0] = [6]int{1, 1, 1, 1, 1, 1}
	for i := 1; i < n; i++ {
		var sum int64 = 0
		for k := 0; k < 6; k++ {
			sum += int64(dp[i-1][k])
		}
		for k := 0; k < 6; k++ {
			var tmpS = sum
			if i-rollMax[k] >= 0 {
				if i-rollMax[k] == 0 {
					tmpS = (tmpS - int64(1) + mod) % mod
				} else {
					for t := 0; t < 6; t++ {
						if t != k {
							tmpS = (tmpS - int64(dp[i-rollMax[k]-1][t]) + mod) % mod
						}
					}
				}
			}
			dp[i][k] = int(tmpS)
		}
	}
	fmt.Println(dp)
	res := 0
	for i := 0; i < 6; i++ {
		res = int(int64(res+dp[len(dp)-1][i]) % mod)
	}
	return res
}
