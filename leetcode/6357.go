/**
 * @author: baowj
 * @time: 2023/2/12 17:48:15
 * @email: bwj_678@qq.com
 */
package main

import "math"

func minimumScore(s string, t string) int {
	leftdp, rightdp := make([]int, len(t)), make([]int, len(t))
	sp, tp := 0, 0
	for tp < len(t) {
		for sp < len(s) && s[sp] != t[tp] {
			sp++
		}
		leftdp[tp] = sp
		if sp != len(s) {
			sp++
		}
		tp++
	}

	for sp, tp = len(s)-1, len(t)-1; tp >= 0; tp-- {
		for sp >= 0 && s[sp] != t[tp] {
			sp--
		}
		rightdp[tp] = sp
		if sp != -1 {
			sp--
		}
	}

	res := math.MaxInt
	if rightdp[0] >= 0 {
		return 0
	}
	l, r := -1, 0
	for l < r && r < len(t) {
		lp := -1
		if l >= 0 {
			lp = leftdp[l]
		}
		rp := rightdp[r]
		if rp > lp {
			if r-l-1 < res {
				res = r - l - 1
			}
			l++
		} else {
			r++
		}
	}
	if l < 0 {
		l = 0
	}
	for l < len(t) && leftdp[l] < len(s) {
		if len(t)-l-1 < res {
			res = len(t) - l - 1
		}
		l++
	}
	if res == math.MaxInt {
		return len(t)
	}
	return res
}
