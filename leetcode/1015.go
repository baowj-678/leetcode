/**
 * @author: baowj
 * @time: 2023/5/10 10:07
 */
package main

func smallestRepunitDivByK(k int) int {
	res := 1
	if res%k == 0 {
		return 1
	}
	for i := 2; i <= k; i++ {
		res = (res*10 + 1) % k
		if res == 0 {
			return i
		}
	}
	return -1
}
