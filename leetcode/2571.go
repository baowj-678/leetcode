/**
 * @author: baowj
 * @time: 2023/2/25 19:53:57
 * @email: bwj_678@qq.com
 */
package main

func minOperations(n int) int {
	res := 0
	continueOne := 0
	for n > 0 {
		if n&0b1 == 0b1 {
			continueOne++
		} else {
			if continueOne == 1 {
				res++
				continueOne = 0
			} else if continueOne >= 2 {
				res++
				n += 0b1
				continueOne = 1
			}
		}
		n >>= 1
	}
	if continueOne == 1 {
		res++
	} else if continueOne >= 2 {
		res += 2
	}
	return res
}
