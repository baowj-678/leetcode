/**
 * @author: baowj
 * @time: 2022/11/22 23:01:06
 * @email: bwj_678@qq.com
 */
package main

var MOD878 int64 = 1e9 + 7

func nthMagicalNumber(n int, a int, b int) int {
	l, r := int64(0), int64(1e14)
	for l < r {
		mid := (l + r) / 2
		tmp := f878(mid, int64(a), int64(b))
		if tmp < int64(n) {
			l = mid + 1
		} else if tmp > int64(n) {
			r = mid - 1
		} else {
			r = mid
		}
	}
	return int(l % MOD878)
}

func f878(x, a, b int64) int64 {
	return x/a + x/b - x/(a*b/lcm878(a, b))
}

func lcm878(x, y int64) int64 {
	if x < y {
		x, y = y, x
	}
	for y != 0 {
		y, x = x%y, y
	}
	return x
}

func main() {
	nthMagicalNumber(1, 2, 3)
}
