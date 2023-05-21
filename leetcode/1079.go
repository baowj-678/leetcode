/**
 * @author: baowj
 * @time: 2023/5/21 15:21:59
 * @email: bwj_678@qq.com
 */
package main

func numTilePossibilities(tiles string) int {
	m := map[int32]int{}
	for _, tile := range tiles {
		if _, ok := m[tile]; !ok {
			m[tile] = 0
		}
		m[tile]++
	}
	res := dfs1079(m) - 1
	return res
}

func dfs1079(m map[int32]int) int {
	res := 1
	for k := range m {
		if m[k] > 0 {
			m[k]--
			res += dfs1079(m)
			m[k]++
		}
	}
	return res
}
