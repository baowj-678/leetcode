/**
 * @author: baowj
 * @time: 2023/3/22 18:54
 */
package main

import (
	"fmt"
	"sort"
)

func bestTeamScore(scores []int, ages []int) int {
	scoreAge := make([][2]int, len(scores))
	for i := 0; i < len(scores); i++ {
		scoreAge[i] = [2]int{scores[i], ages[i]}
	}
	max := func(a, b int) int {
		if a > b {
			return a
		}
		return b
	}
	sort.Slice(scoreAge, func(i, j int) bool {
		if scoreAge[i][0] == scoreAge[j][0] {
			return scoreAge[i][1] <= scoreAge[j][1]
		} else {
			return scoreAge[i][0] < scoreAge[j][0]
		}
	})
	fmt.Println(scoreAge)
	dp := make([]int, len(scoreAge))
	for i := 0; i < len(dp); i++ {
		score, age := scoreAge[i][0], scoreAge[i][1]
		for j := 0; j < i; j++ {
			_, aj := scoreAge[j][0], scoreAge[j][1]
			if aj <= age {
				dp[i] = max(dp[i], dp[j])
			}
		}
		dp[i] += score
	}
	fmt.Println(dp)
	res := 0
	for i := 0; i < len(dp); i++ {
		res = max(res, dp[i])
	}
	return res
}
