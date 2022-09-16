/**
 * @author: baowj
 * @time: 2022/8/26 21:06:51
 */
package main

func isPowerOfTwo(n int) bool {
	cnt := 0
	for n > 0 {
		if n&0x1 == 0x1 {
			cnt++
		}
		n >>= 1
	}
	if cnt == 1 {
		return true
	}
	return false
}
