/**
 * @author: baowj
 * @time: 2023/8/30 10:11
 */
package main

func minimumJumps(forbidden []int, a int, b int, x int) int {
	maxF := 0
	forbiddenSet := map[int]bool{}
	max := func(x, y int) int {
		if x < y {
			x = y
		}
		return x
	}
	min := func(x, y int) int {
		if x > y {
			x = y
		}
		return x
	}
	for _, num := range forbidden {
		maxF = max(maxF, num)
		forbiddenSet[num] = true
	}

	var bound int
	if a > b {
		bound = x + b
	} else if a == b {
		bound = x
	} else {
		bound = max(maxF+a+b, x)
	}
	dp := make([][2]int, bound+1) // -1: unknown, not visited, 0+: known, -2: visited, -3: can't reach
	for i := range dp {
		dp[i] = [2]int{-1, -1}
	}
	dp[0][0], dp[0][1] = 0, 0
	var bfs func(p int, direct int) int // return: 0+: min steps, -1: can't reach

	bfs = func(p int, direct int) int {
		// 0 right, 1 left
		if p > bound || p < 0 || forbiddenSet[p] {
			return -1
		}
		if dp[p][direct] >= 0 {
			return dp[p][direct]
		} else if dp[p][direct] == -2 || dp[p][direct] == -3 {
			return -1
		} else {
			dp[p][direct] = -2
		}

		if direct == 0 {
			r := bfs(p-a, 0)
			l := bfs(p-a, 1)
			if r == -1 && l == -1 {
				dp[p][direct] = -3
				return -1
			}
			if r == -1 {
				dp[p][direct] = l + 1
				return l + 1
			}
			if l == -1 {
				dp[p][direct] = r + 1
				return r + 1
			}
			dp[p][direct] = min(l, r) + 1
			return min(l, r) + 1
		} else {
			r := bfs(p+b, 0)
			if r == -1 {
				dp[p][direct] = -3
				return -1
			} else {
				dp[p][direct] = r + 1
				return r + 1
			}
		}
	}
	r := bfs(x, 0)
	l := bfs(x, 1)
	if r == -1 && l == -1 {
		return -1
	}
	if r == -1 {
		return l
	}
	if l == -1 {
		return r
	}
	return min(l, r)
}
