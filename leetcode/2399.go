/**
 * @author: baowj
 * @time: 2022/9/17 20:21:32
 */
package main

func checkDistances(s string, distance []int) bool {
	firstIdx := make([]int, 26)
	for i := 0; i < 26; i++ {
		firstIdx[i] = -1
	}
	for i, c := range s {
		c -= 'a'
		if firstIdx[c] == -1 {
			firstIdx[c] = i
		} else {
			if distance[c] != i-firstIdx[c]-1 {
				return false
			} else {
				distance[c] = 0
			}
		}
	}
	return true
}
