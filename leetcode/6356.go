/**
 * @author: baowj
 * @time: 2023/2/12 17:13:12
 * @email: bwj_678@qq.com
 */
package main

func substringXorQueries(s string, queries [][]int) [][]int {
	queryvals := make([]int, len(queries))
	m := map[int][]int{}
	for i := 0; i < len(queries); i++ {
		queryvals[i] = queries[i][0] ^ queries[i][1]
		m[queryvals[i]] = []int{-1, -1}
	}
	res := make([][]int, len(queries))
	for l := 0; l < len(s); l++ {
		if s[l] == '0' {
			if v, ok := m[0]; ok {
				if v[0] == -1 {
					m[0] = []int{l, l}
				}
			}
			continue
		}
		r := l + 29
		if r >= len(s) {
			r = len(s) - 1
		}
		x := 0
		for i := l; i <= r; i++ {
			x *= 2
			if s[i] == '1' {
				x = x + 1
			}
			if v, ok := m[x]; ok {
				if v[0] == -1 {
					m[x] = []int{l, i}
				}
			}
		}
	}

	for i := 0; i < len(queryvals); i++ {
		res[i] = m[queryvals[i]]
	}
	return res
}
