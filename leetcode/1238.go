/**
 * @author: baowj
 * @time: 2023/2/23 9:55
 */
package main

func circularPermutation(n int, start int) []int {
	res := make([]int, 1<<n)
	res[1] = 1
	for k := 1; k < n; k++ {
		for i, j := 1<<k, (1<<k)-1; i < 1<<(k+1); i++ {
			res[i] = res[j] | (1 << k)
			j--
		}
	}

	for i := 0; i < len(res); i++ {
		if res[i] == start {
			res = append(res[i:], res[:i]...)
		}
	}
	return res
}
