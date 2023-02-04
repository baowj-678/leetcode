/**
 * @author: baowj
 * @time: 2023/2/4 21:19:16
 * @email: bwj_678@qq.com
 */
package main

func monkeyMove(n int) int {
	const MOD2550 int64 = 1e9 + 7
	x, exp := int64(1), int64(2)
	for n > 0 {
		if n&0x1 == 1 {
			x = (x * exp) % MOD2550
		}
		exp = (exp * exp) % MOD2550
		n >>= 1
	}
	return int((x - 2 + MOD2550) % MOD2550)
}
