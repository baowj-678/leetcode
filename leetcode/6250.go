/**
 * @author: baowj
 * @time: 2022/11/27 16:04:40
 * @email: bwj_678@qq.com
 */
package main

import "math"

func bestClosingTime(customers string) int {
	yCnt := make([]int, len(customers))
	last := 0
	for i := len(customers) - 1; i >= 0; i-- {
		if customers[i] == 'Y' {
			last++
		}
		yCnt[i] = last
	}
	last = 0
	minCost := math.MaxInt
	time := 0
	for i, c := range customers {
		cost := last + yCnt[i]
		if cost < minCost {
			time = i
			minCost = cost
		}
		if c == 'N' {
			last++
		}
	}
	if last < minCost {
		return len(customers)
	}
	return time
}

func main() {
	println(bestClosingTime("YYYY"))
}
