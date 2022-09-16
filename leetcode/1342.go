/**
 * @author: baowj
 * @time: 2022/8/26 10:14:12
 */
package main

func numberOfSteps(num int) int {
	cnt := 0
	for num > 1 {
		if num&0x1 > 0 {
			cnt += 2
		} else {
			cnt++
		}
		num >>= 1
	}
	if num == 1 {
		cnt++
	}
	return cnt
}

func main() {
	numberOfSteps(14)
}
