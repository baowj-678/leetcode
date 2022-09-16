/**
 * @author: baowj
 * @time: 2022/8/21 10:31:50
 */
package main

func minimumRecolors(blocks string, k int) int {
	blackCnt, maxBlackCnt := 0, 0
	for i, c := range blocks {
		if i-k >= 0 {
			if blocks[i-k] == 'B' {
				blackCnt--
			}
		}
		if c == 'B' {
			blackCnt++
		}
		if blackCnt > maxBlackCnt {
			maxBlackCnt = blackCnt
		}
	}
	return k - maxBlackCnt
}
