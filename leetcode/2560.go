/**
 * @author: baowj
 * @time: 2023/2/6 22:13:58
 * @email: bwj_678@qq.com
 */
package main

import "sort"

func minCapability(nums []int, k int) int {
	check := func(mx int) bool {
		f0, f1 := 0, 0
		for _, num := range nums {
			if num > mx {
				f0 = f1
			} else {
				if f0+1 > f1 {
					f1, f0 = f0+1, f1
				} else {
					f0 = f1
				}
			}
		}
		return f1 >= k
	}
	mx := 0
	for _, num := range nums {
		if num > mx {
			mx = num
		}
	}
	res := sort.Search(mx+1, func(i int) bool {
		return check(i)
	})
	return res
}
