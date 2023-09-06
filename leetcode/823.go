/**
 * @author: baowj
 * @time: 2023/8/29 16:17
 */
package main

import "sort"

func numFactoredBinaryTrees(arr []int) int {
	const MaxValue = int64(1e9 + 7)
	existNum := map[int]bool{}
	sort.Ints(arr)
	for _, num := range arr {
		existNum[num] = true
	}
	numSubs := map[int][][2]int{}
	for _, num := range arr {
		for i := 0; i < len(arr); i++ {
			x := arr[i]
			y := num / x
			if x > y {
				break
			}
			if x*y == num && existNum[y] {
				if subs, ok := numSubs[num]; ok {
					numSubs[num] = append(subs, [2]int{x, y})
				} else {
					numSubs[num] = [][2]int{{x, y}}
				}
			}
		}
	}
	fnTable := map[int]int64{}

	var fn func(value int) int64
	fn = func(value int) int64 {
		if res, ok := fnTable[value]; ok {
			return res
		}
		res := int64(1)
		subNums := numSubs[value]
		if subNums == nil {
			return res
		}

		for _, subNum := range subNums {
			if subNum[0] == subNum[1] {
				v := fn(subNum[0])
				res = (res + ((v * v) % MaxValue)) % MaxValue
			} else {
				vL := fn(subNum[0])
				vR := fn(subNum[1])
				res = (res + ((2 * ((vL * vR) % MaxValue)) % MaxValue)) % MaxValue
			}
		}
		res %= MaxValue
		fnTable[value] = res
		return res
	}

	res := int64(0)
	for _, num := range arr {
		res = (res + fn(num)) % MaxValue
	}
	return int(res)
}
