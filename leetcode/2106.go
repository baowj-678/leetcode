/**
 * @author: baowj
 * @time: 2023/5/5 15:50
 */
package main

func maxTotalFruits(fruits [][]int, startPos int, k int) int {
	maxLen := fruits[len(fruits)-1][0]
	if maxLen < startPos {
		maxLen = startPos
	}
	maxLen++
	fruitSum := make([]int, maxLen)
	for _, fruit := range fruits {
		fruitSum[fruit[0]] = fruit[1]
	}
	sum := 0
	for i := 0; i < len(fruitSum); i++ {
		fruitSum[i] += sum
		sum = fruitSum[i]
	}

	left := startPos - k
	if left < 0 {
		left = 0
	}

	max := 0
	for i := left; i <= startPos; i++ {
		right := k - (startPos-i)*2 + startPos
		if right < startPos {
			right = startPos
		}
		if right > maxLen-1 {
			right = maxLen - 1
		}

		if i > 0 {
			sum = fruitSum[right] - fruitSum[i-1]
		} else {
			sum = fruitSum[right]
		}
		if sum > max {
			max = sum
		}
	}

	right := startPos + k
	if right > maxLen-1 {
		right = maxLen - 1
	}

	for i := startPos + 1; i <= right; i++ {
		left = startPos - (k - (i-startPos)*2)
		if left > startPos {
			left = startPos
		}
		if left < 0 {
			left = 0
		}

		if left > 0 {
			sum = fruitSum[i] - fruitSum[left-1]
		} else {
			sum = fruitSum[i]
		}
		if sum > max {
			max = sum
		}
	}
	return max
}
