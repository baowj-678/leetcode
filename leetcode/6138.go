/**
 * @author: baowj
 * @time: 2022/8/7 17:11:52
 */
package main

func longestIdealString(s string, k int) int {
	dp := make([]int, 26)
	for i := 0; i < 26; i++ {
		dp[i] = 0
	}
	for _, c := range s {
		p := int(c) - int('a')
		start := p - k
		end := p + k
		if start < 0 {
			start = 0
		}
		if end > 25 {
			end = 25
		}
		max := 0
		for i := start; i <= end; i++ {
			if max < dp[i] {
				max = dp[i]
			}
		}
		if dp[p] < max {
			dp[p] = max + 1
		} else {
			dp[p] += 1
		}
	}
	res := 0
	for i := 0; i < 26; i++ {
		if res < dp[i] {
			res = dp[i]
		}
	}
	return res
}

func main() {
	s := "fbv"
	println(longestIdealString(s, 2))
}
