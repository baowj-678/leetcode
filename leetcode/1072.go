/**
 * @author: baowj
 * @time: 2023/5/17 11:04
 */
package main

func maxEqualRowsAfterFlips(matrix [][]int) int {
	m := map[string]int{}
	for _, row := range matrix {
		s := ""
		isZero := row[0] == 1
		for _, c := range row {
			if isZero {
				switch c {
				case 0:
					s += "0"
				case 1:
					s += "1"
				}
			} else {
				switch c {
				case 0:
					s += "1"
				case 1:
					s += "0"
				}
			}
		}
		if _, ok := m[s]; !ok {
			m[s] = 0
		}
		m[s]++
	}

	res := 1
	for _, v := range m {
		if v > res {
			res = v
		}
	}
	return res
}
