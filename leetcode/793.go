/**
 * @author: baowj
 * @time: 2022/8/28 20:52:51
 */
package main

func zeta(n int) (res int) {
	for n > 0 {
		n /= 5
		res += n
	}
	return
}

func nx(k int) int {
	l, r := 0, 5*k
	for l < r {
		m := (l + r) / 2
		if zeta(m) < k {
			l = m + 1
		} else {
			r = m
		}
	}
	return l
}

func preimageSizeFZF(k int) int {
	return nx(k+1) - nx(k)
}
