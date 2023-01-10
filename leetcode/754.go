/**
 * @author: baowj
 * @time: 2022/11/6 19:01:33
 */
package main

func reachNumber(target int) int {
	res := 0
	p := 0
	if target < 0 {
		target = -target
	}
	for {
		res++
		p += res
		if p >= target {
			break
		}
	}
	if p == target {
		return res
	} else if (p-target)%2 == 0 {
		return res
	} else {
		if res%2 == 0 {
			return res + 1
		} else {
			return res + 2
		}
	}
}
