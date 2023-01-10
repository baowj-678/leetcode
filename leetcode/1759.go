/**
 * @author: baowj
 * @time: 2022/12/27 21:25:37
 * @email: bwj_678@qq.com
 */
package main

const MOD1759 int = 1e9 + 7

func countHomogenous(s string) int {
	last := 0
	res := 0
	for i := 0; i < len(s); i++ {
		if i-1 >= 0 && s[i] != s[i-1] {
			n := i - last
			res = (res + ((n+1)*n/2)%MOD1759) % MOD1759
			last = i
		}
	}
	n := len(s) - last
	res = (res + ((n+1)*n/2)%MOD1759) % MOD1759

	return res
}
