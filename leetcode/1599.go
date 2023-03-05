/**
 * @author: baowj
 * @time: 2023/3/5 10:25:08
 * @email: bwj_678@qq.com
 */
package main

func minOperationsMaxProfit(customers []int, boardingCost int, runningCost int) int {
	maxProfit, maxOp := 0, -1
	tmp := 0
	csum := 0
	for i := 0; i < len(customers) || csum > 0; i++ {
		if i < len(customers) {
			customer := customers[i]
			csum += customer
		}

		if csum >= 4 {
			csum -= 4
			tmp += 4*boardingCost - runningCost
		} else {
			tmp += csum*boardingCost - runningCost
			csum = 0
		}
		if tmp > maxProfit {
			maxProfit = tmp
			maxOp = i + 1
		}
	}
	return maxOp
}
