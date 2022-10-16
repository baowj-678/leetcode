/**
 * @author: baowj
 * @time: 2022/9/17 14:39:53
 */
package main

func maxLengthBetweenEqualCharacters(s string) int {
	firstIdx := make([]int, 26)
	for i := 0; i < len(firstIdx); i++ {
		firstIdx[i] = -1
	}
	res := -1
	for i, c := range s {
		c -= 'a'
		if firstIdx[c] == -1 {
			firstIdx[c] = i
		} else {
			if i-firstIdx[c]-1 > res {
				res = i - firstIdx[c] - 1
			}
		}
	}
	return res
}
