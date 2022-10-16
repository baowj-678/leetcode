/**
 * @author: baowj
 * @time: 2022/9/18 11:21:51
 */
package main

func minimumMoney(transactions [][]int) int64 {
	var res int64 = 0
	var sum int64 = 0
	for _, transaction := range transactions {
		if transaction[0]-transaction[1] > 0 {
			sum += int64(transaction[0] - transaction[1])
		}
	}
	for _, transaction := range transactions {
		if transaction[1] >= transaction[0] {
			if res < int64(transaction[0])+sum {
				res = int64(transaction[0]) + sum
			}
		} else {
			if res < sum-int64(transaction[0]-transaction[1])+int64(transaction[0]) {
				res = sum - int64(transaction[0]-transaction[1]) + int64(transaction[0])
			}
		}
	}
	return res
}
