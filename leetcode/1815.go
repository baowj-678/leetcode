/**
 * @author: baowj
 * @time: 2023/1/22 20:55:55
 * @email: bwj_678@qq.com
 */
package main

import "sort"

func maxHappyGroups(batchSize int, groups []int) int {
	for i, group := range groups {
		groups[i] = group % batchSize
	}
	sort.Ints(groups)
	res := 0
	i := 0
	for i < len(groups) && groups[i] == 0 {
		i++
	}
	res += i
	groups = groups[i:]

	dp := map[int64]int{}
	dpMod := map[int64]int{}
	dp[0] = 0
	dpMod[0] = 0

	mods := make([]int, batchSize)
	for i := 0; i < len(groups); i++ {
		mods[groups[i]]++
	}
	mods[0] = 0
	state := int64(0)
	for i, mod := range mods {
		state |= int64(mod) << (5 * int64(i))
	}
	res += dfs1815(dp, dpMod, state, batchSize)

	return res
}

func dfs1815(dp, dpMod map[int64]int, state int64, batchSize int) int {
	if v, ok := dp[state]; ok {
		return v
	}
	res := 0
	for i := 0; i < batchSize; i++ {
		num := stateToNum(state, i)
		if num > 0 {
			newState := setStateNum(state, i, num-1)
			v := dfs1815(dp, dpMod, newState, batchSize)

			if _, ok := dpMod[state]; !ok {
				dpMod[state] = (dpMod[newState] + i) % batchSize
			}
			mod := dpMod[newState]
			if mod == 0 {
				if res < v+1 {
					res = v + 1
				}
			} else {
				if res < v {
					res = v
				}
			}
		}
	}
	dp[state] = res
	return res
}

func stateToNum(state int64, idx int) int {
	return int((state >> (5 * idx)) & 0b11111)
}

func setStateNum(state int64, idx int, v int) int64 {
	return (state & (^(int64(0b11111) << (5 * idx)))) | (int64(v) << (5 * idx))
}
