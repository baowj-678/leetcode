/**
 * @author: baowj
 * @time: 2022/9/5 10:32:33
 */
package main

func finalPrices(prices []int) []int {
	for i, x := range prices {
		for j := i + 1; j < len(prices); j++ {
			if prices[j] <= x {
				prices[i] -= prices[j]
				break
			}
		}
	}
	return prices
}
