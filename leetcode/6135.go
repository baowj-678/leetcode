/**
 * @author: baowj
 * @time: 2022/7/31 11:06:43
 */
package main

func longestCycle(edges []int) int {
	flag := make([]int, len(edges))
	for i := 0; i < len(flag); i++ {
		flag[i] = -1
	}
	maxLen := 0
	tmpLen := 0
	flagV := 1
	for i := 0; i < len(edges); i++ {
		tmp := i
		for flag[tmp] == -1 {
			flag[tmp] = flagV
			tmp = edges[tmp]
			if tmp == -1 {
				break
			}
		}
		if tmp != -1 {
			if flag[tmp] == flagV {
				start := tmp
				tmp = edges[tmp]
				tmpLen = 1
				for start != tmp {
					tmp = edges[tmp]
					tmpLen += 1
				}
				if tmpLen > maxLen {
					maxLen = tmpLen
				}
			}
		}
		flagV += 1
	}
	if maxLen == 0 {
		return -1
	}
	return maxLen
}
