/**
 * @author: baowj
 * @time: 2022/4/14 09:37:11
 */
package main

func maximumWealth(accounts [][]int) int {
	max := 0
	for _, list := range accounts {
		tmp := 0
		for _, wealth := range list {
			tmp += wealth
		}
		if tmp > max {
			max = tmp
		}
	}
	return max
}
