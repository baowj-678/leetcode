/**
 * @author: baowj
 * @time: 2022/9/6 10:01:06
 */
package main

func uniqueLetterString(s string) int {
	idx := [26][]int{}
	for i, c := range s {
		idx[c-'a'] = append(idx[c-'a'], i)
	}
	res := 0
	for i := 0; i < 26; i++ {
		for j := 0; j < len(idx[i]); j++ {
			a, b := 0, 0
			if j == 0 {
				a = idx[i][j] + 1
			} else {
				a = idx[i][j] - idx[i][j-1]
			}

			if j == len(idx[i])-1 {
				b = len(s) - idx[i][j]
			} else {
				b = idx[i][j+1] - idx[i][j]
			}
			res += a * b
		}
	}
	return res
}
