/**
 * @author: baowj
 * @time: 2023/6/1 20:14
 */
package main

import (
	"sort"
)

func maximumTastiness(price []int, k int) int {
	sort.Ints(price)
	check := func(sweet int) bool {
		cnt := 0
		for i, last := 0, -int(1e9); i < len(price); i++ {
			if price[i]-last >= sweet {
				last = price[i]
				cnt++
			}
			if cnt >= k {
				return true
			}
		}
		return false
	}
	l, r := 0, int(1e9)
	for l < r {
		mid := (l + r + 1) / 2
		if check(mid) {
			l = mid
		} else {
			r = mid - 1
		}
	}
	return l
}
