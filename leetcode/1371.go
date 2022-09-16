/**
 * @author: baowj
 * @time: 2022/9/5 11:20:50
 */
package main

func findTheLongestSubstring(s string) int {
	flag := 0
	table := make([]int, 32)
	for i := 0; i < len(table); i++ {
		table[i] = -2
	}
	str := "aeiou"
	maxLen := 0
	table[0] = -1
	for i, c := range s {
		for j, t := range str {
			if t == c {
				flag ^= 1 << j
			}
		}
		if table[flag] == -2 {
			table[flag] = i
		}
		if i-table[flag] > maxLen {
			maxLen = i - table[flag]
		}
	}
	return maxLen
}
