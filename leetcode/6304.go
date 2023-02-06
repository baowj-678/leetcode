/**
 * @author: baowj
 * @time: 2023/2/5 16:27:00
 * @email: bwj_678@qq.com
 */
package main

func maxCount(banned []int, n int, maxSum int) int {
	res := 0
	bannedSet := map[int]struct{}{}
	for _, x := range banned {
		bannedSet[x] = struct{}{}
	}
	sum := 0
	for i := 1; i <= n; i++ {
		if _, ok := bannedSet[i]; !ok {
			sum += i
			res++
			if sum > maxSum {
				return res - 1
			}
		}
	}
	return res
}
