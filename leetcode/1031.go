/**
 * @author: baowj
 * @time: 2023/4/26 09:35:13
 * @email: bwj_678@qq.com
 */
package main

func maxSumTwoNoOverlap(nums []int, firstLen int, secondLen int) int {
	prefixSum := make([]int, len(nums))
	s := 0
	for i, n := range nums {
		s += n
		prefixSum[i] = s
	}
	max := func(a, b int) int {
		if a < b {
			a = b
		}
		return a
	}
	res := 0
	// [i, i + first)
	for i := 0; i+firstLen+secondLen <= len(nums); i++ {
		for j := i + firstLen; j+secondLen <= len(nums); j++ {
			if i > 0 {
				res = max(res, prefixSum[i+firstLen-1]-prefixSum[i-1]+prefixSum[j+secondLen-1]-prefixSum[j-1])
			} else {
				res = max(res, prefixSum[i+firstLen-1]+prefixSum[j+secondLen-1]-prefixSum[j-1])
			}
		}
	}
	firstLen, secondLen = secondLen, firstLen
	for i := 0; i+firstLen+secondLen <= len(nums); i++ {
		for j := i + firstLen; j+secondLen <= len(nums); j++ {
			if i > 0 {
				res = max(res, prefixSum[i+firstLen-1]-prefixSum[i-1]+prefixSum[j+secondLen-1]-prefixSum[j-1])
			} else {
				res = max(res, prefixSum[i+firstLen-1]+prefixSum[j+secondLen-1]-prefixSum[j-1])
			}
		}
	}
	return res
}
