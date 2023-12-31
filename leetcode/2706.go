package main

import "math"

func buyChoco(prices []int, money int) int {
	minMoneys := [2]int{math.MaxInt, math.MaxInt}
	for _, price := range prices {
		if price < minMoneys[0] {
			minMoneys[1], minMoneys[0] = minMoneys[0], price
			continue
		}
		if price < minMoneys[1] {
			minMoneys[1] = price
		}
	}
	if minMoneys[0]+minMoneys[1] > money {
		return money
	}
	return money - minMoneys[0] - minMoneys[1]
}
