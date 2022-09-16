/**
 * @author: baowj
 * @time: 2022/8/13 10:12:59
 */
package main

func maxChunksToSorted2(arr []int) int {
	stack := []int{}
	for _, x := range arr {
		if len(stack) > 0 {
			if x >= stack[len(stack)-1] {
				stack = append(stack, x)
			} else {
				max := stack[len(stack)-1]
				for len(stack) > 0 && stack[len(stack)-1] > x {
					stack = stack[:len(stack)-1]
				}
				stack = append(stack, max)
			}
		} else {
			stack = append(stack, x)
		}
	}
	return len(stack)
}

func main() {
	arr := []int{5, 4, 3, 2, 1}
	println(maxChunksToSorted2(arr))
}
