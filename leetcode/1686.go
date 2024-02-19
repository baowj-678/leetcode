package main

import "sort"

func stoneGameVI(aliceValues []int, bobValues []int) int {
	sumValue := make([][]int, 0, len(aliceValues))
	for i := 0; i < len(aliceValues); i++ {
		sumValue = append(sumValue, []int{aliceValues[i] + bobValues[i], i})
	}
	sort.Slice(sumValue, func(i, j int) bool {
		return sumValue[i][0] >= sumValue[j][0]
	})
	a, b := 0, 0
	isAlice := true
	for _, value := range sumValue {
		if isAlice {
			a += aliceValues[value[1]]
		} else {
			b += bobValues[value[1]]

		}
		isAlice = !isAlice
	}
	if a > b {
		return 1
	} else if a < b {
		return -1
	}
	return 0
}
