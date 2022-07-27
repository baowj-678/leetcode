/**
 * @author: baowj
 * @time: 2022/4/24 09:50:54
 */
package main

import "math"

func binaryGap(n int) int {
	gap := 0
	for cnt := math.MinInt; n > 0; {
		cnt++
		if n&1 == 1 {
			if cnt > gap {
				gap = cnt
			}
			cnt = 0
		}
		n >>= 1
	}
	return gap
}

func main() {
	println(binaryGap(22))
	//print(22&1 == 1)
}
