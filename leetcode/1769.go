/**
 * @author: baowj
 * @time: 2022/12/3 15:36:28
 * @email: bwj_678@qq.com
 */
package main

func minOperations1769(boxes string) []int {
	res := make([]int, len(boxes))
	opCnt, ballCnt := 0, 0
	for i, box := range boxes {
		opCnt += ballCnt
		res[i] += opCnt
		if box == '1' {
			ballCnt++
		}
	}
	opCnt, ballCnt = 0, 0
	for i := len(boxes) - 1; i >= 0; i-- {
		opCnt += ballCnt
		res[i] += opCnt
		if boxes[i] == '1' {
			ballCnt++
		}
	}
	return res
}
