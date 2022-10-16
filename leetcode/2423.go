/**
 * @author: baowj
 * @time: 2022/10/8 23:54:02
 */
package main

import "math"

func equalFrequency(word string) bool {
	cnt := make([]int, 26)
	for _, c := range word {
		cnt[c-'a']++
	}
	for _, c := range word {
		cnt[c-'a']--
		min, max := math.MaxInt, 0
		for _, x := range cnt {
			if x > 0 && min > x {
				min = x
			}
			if x > max {
				max = x
			}
		}
		if max == min {
			return true
		}
		cnt[c-'a']++
	}
	return false
}
