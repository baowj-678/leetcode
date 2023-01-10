/**
 * @author: baowj
 * @time: 2022/11/2 16:59:55
 */
package main

func averageValue(nums []int) int {
	sum, cnt := 0, 0
	for _, num := range nums {
		if num%2 == 0 && num%3 == 0 {
			sum += num
			cnt++
		}
	}
	if cnt == 0 {
		return 0
	}
	return sum / cnt
}
