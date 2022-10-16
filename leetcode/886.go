/**
 * @author: baowj
 * @time: 2022/10/16 18:40:47
 */
package main

func possibleBipartition(n int, dislikes [][]int) bool {
	nextM := make([][]int, n+1)
	for i := 0; i < len(nextM); i++ {
		nextM[i] = []int{}
	}
	for _, pair := range dislikes {
		nextM[pair[0]] = append(nextM[pair[0]], pair[1])
		nextM[pair[1]] = append(nextM[pair[1]], pair[0])
	}
	partition := make([]int, n+1)
	stack := []int{}
	for i := 1; i <= n; i++ {
		if partition[i] == 0 {
			stack = append(stack, i)
			for len(stack) > 0 {
				first := stack[0]
				stack = stack[1:]
				if partition[first] == 0 {
					partition[first] = 1
					stack = append(stack, first)
				}
				if partition[first] == 1 {
					for _, next := range nextM[first] {
						if partition[next] == 1 {
							return false
						} else if partition[next] == 0 {
							partition[next] = 2
							stack = append(stack, next)
						}
					}
				}
				if partition[first] == 2 {
					for _, next := range nextM[first] {
						if partition[next] == 2 {
							return false
						} else if partition[next] == 0 {
							partition[next] = 1
							stack = append(stack, next)
						}
					}
				}
			}
		}
	}
	return true
}
