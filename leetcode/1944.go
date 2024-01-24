package main

func canSeePersonsCount(heights []int) []int {
	res := make([]int, len(heights))
	stack := [][2]int{}
	for i := len(heights) - 1; i >= 0; i-- {
		height := heights[i]
		for len(stack) > 0 && stack[len(stack)-1][0] < height {
			res[i]++
			stack = stack[:len(stack)-1]
		}
		if len(stack) > 0 {
			res[i]++
		}
		stack = append(stack, [2]int{height, i})
	}
	return res
}
