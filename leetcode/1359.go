/**
 * @author: baowj
 * @time: 2022/8/31 21:49:48
 */
package main

const MOD = 1000000000 + 7

func countOrders(n int) int {
	var res int64 = 1
	var i int64 = 1
	for ; i <= int64(2*n); i++ {
		if i%2 == 0 {
			res = (res * i / 2) % MOD
		} else {
			res = (res * i) % MOD
		}
	}
	return int(res)
}
