/**
 * @author: baowj
 * @time: 2022/12/18 18:19:48
 * @email: bwj_678@qq.com
 */
package main

func cycleLengthQueries(n int, queries [][]int) []int {
	res := make([]int, len(queries))
	for i, query := range queries {
		for p := 0; p <= 30; p++ {
			u := query[0] >> p
			find := false
			if u == 0 {
				break
			}
			for q := 0; q <= 30; q++ {
				v := query[1] >> q
				if v == 0 {
					break
				}
				if u == v {
					res[i] = p + q + 1
					find = true
					break
				}
			}
			if find {
				break
			}
		}
	}
	return res
}
