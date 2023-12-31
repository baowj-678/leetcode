package main

import "fmt"

func maximumSumOfHeights(maxHeights []int) int64 {
	prefix := make([]int64, len(maxHeights))
	suffix := make([]int64, len(maxHeights))
	stack := [][2]int{}
	for i, height := range maxHeights {
		for len(stack) > 0 && stack[len(stack)-1][1] > height {
			stack = stack[:len(stack)-1]
		}
		if len(stack) > 0 {
			lastIdx := stack[len(stack)-1][0]
			prefix[i] = prefix[lastIdx] + int64(height)*int64(i-lastIdx)
		} else {
			prefix[i] = int64(height) * int64(i+1)
		}
		stack = append(stack, [2]int{i, height})
	}

	stack = [][2]int{}
	for i := len(maxHeights) - 1; i >= 0; i-- {
		height := maxHeights[i]
		for len(stack) > 0 && stack[len(stack)-1][1] > height {
			stack = stack[:len(stack)-1]
		}
		if len(stack) > 0 {
			lastIdx := stack[len(stack)-1][0]
			suffix[i] = suffix[lastIdx] + int64(height)*int64(lastIdx-i)
		} else {
			suffix[i] = int64(height) * int64(len(maxHeights)-i)
		}
		stack = append(stack, [2]int{i, height})
	}
	res := int64(0)
	for i := 0; i < len(maxHeights); i++ {
		if prefix[i]+suffix[i]-int64(maxHeights[i]) > res {
			res = prefix[i] + suffix[i] - int64(maxHeights[i])
		}
	}
	return res
}

func main() {
	nums := []int{5, 3, 4, 1, 1}
	fmt.Println(maximumSumOfHeights(nums))
}
