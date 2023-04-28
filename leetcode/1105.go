/**
 * @author: baowj
 * @time: 2023/4/23 10:05
 */
package main

import "math"

func minHeightShelves(books [][]int, shelfWidth int) int {
	dp := make([]int, len(books))
	// totalHeight
	max := func(x, y int) int {
		if x < y {
			x = y
		}
		return x
	}
	for i, _ := range books {
		currWidth, currHeight := 0, 0
		totalHeight := math.MaxInt
		for j := i; j >= 0; j-- {
			width, height := books[j][0], books[j][1]
			currHeight = max(height, currHeight)
			currWidth += width
			if currWidth <= shelfWidth {
				if j-1 >= 0 {
					if currHeight+dp[j-1] < totalHeight {
						totalHeight = currHeight + dp[j-1]
					}
				} else {
					totalHeight = currHeight
				}
			} else {
				break
			}
		}
		dp[i] = totalHeight
	}
	return dp[len(dp)-1]
}
