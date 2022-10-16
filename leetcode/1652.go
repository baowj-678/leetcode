/**
 * @author: baowj
 * @time: 2022/9/24 09:02:30
 */
package main

func decrypt(code []int, k int) []int {
	sums := make([]int, len(code))
	sums[0] = code[0]
	for i := 1; i < len(sums); i++ {
		sums[i] = sums[i-1] + code[i]
	}
	for i := 0; i < len(code); i++ {
		if k > 0 {
			if i+k < len(code) {
				code[i] = sums[i+k] - sums[i]
			} else {
				code[i] = sums[len(sums)-1] - sums[i] + sums[k-(len(sums)-i-1)-1]
			}
		} else if k < 0 {
			if i+k > 0 {
				code[i] = sums[i-1] - sums[i+k-1]
			} else {
				if i-1 >= 0 {
					code[i] = sums[i-1]
				} else {
					code[i] = 0
				}
				code[i] += sums[len(sums)-1] - sums[len(sums)-(-k-i)-1]
			}
		} else {
			code[i] = 0
		}
	}
	return code
}
