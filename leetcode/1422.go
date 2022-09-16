/**
 * @author: baowj
 * @time: 2022/8/14 09:59:53
 */
package main

func maxScore(s string) int {
	leftArray := make([]int, len(s))
	rightArray := make([]int, len(s))
	leftTmp, rightTmp := 0, 0
	for i := 0; i < len(s); i++ {
		j := len(s) - i - 1
		if s[i] == '0' {
			leftTmp += 1
		}
		if s[j] == '1' {
			rightTmp += 1
		}
		leftArray[i] = leftTmp
		rightArray[j] = rightTmp
	}
	max := rightArray[0]
	for i := 0; i < len(s)-1; i++ {
		tmp := leftArray[i] + rightArray[i+1]
		if tmp > max {
			max = tmp
		}
	}
	if leftArray[len(leftArray)-1] > max {
		max = leftArray[len(leftArray)-1]
	}
	return max
}
