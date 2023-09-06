/**
 * @author: baowj
 * @time: 2023/8/31 10:01
 */
package main

func maxDistToClosest(seats []int) int {
	res := 0
	last := -1

	maxFn := func(x, y int) int {
		if x < y {
			x = y
		}
		return x
	}
	for i, seat := range seats {
		if seat == 1 {
			if last == -1 {
				res = maxFn(res, i)
			} else {
				res = maxFn(res, (i-last)/2)
			}
			last = i
		}
	}
	res = maxFn(res, len(seats)-last-1)
	return res
}
