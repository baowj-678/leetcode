/**
 * @author: baowj
 * @time: 2022/9/29 10:38:39
 */
package main

func isInterleave(s1 string, s2 string, s3 string) bool {
	if len(s1)+len(s2) != len(s3) {
		return false
	}
	dp := make([]bool, (len(s1)+1)*(len(s2)+1))
	width := len(s2) + 1
	dp[0] = true
	for _, c := range s3 {
		for i := len(s1); i >= 0; i-- {
			for j := len(s2); j >= 0; j-- {
				if i >= 1 && s1[i-1] == uint8(c) && dp[(i-1)*width+j] {
					dp[i*width+j] = true
				} else if j >= 1 && s2[j-1] == uint8(c) && dp[i*width+j-1] {
					dp[i*width+j] = true
				} else {
					dp[i*width+j] = false
				}
			}
		}
	}
	return dp[len(s1)*width+len(s2)]
}
