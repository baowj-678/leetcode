/**
 * @author: baowj
 * @time: 2023/4/15 10:16:59
 * @email: bwj_678@qq.com
 */
package main

func gardenNoAdj(n int, paths [][]int) []int {
	nextNode := make([][]int, n)
	for i := 0; i < n; i++ {
		nextNode[i] = []int{}
	}
	for _, path := range paths {
		i, j := path[0]-1, path[1]-1
		nextNode[i] = append(nextNode[i], j)
		nextNode[j] = append(nextNode[j], i)
	}
	res := make([]int, n)
	for i := 0; i < n; i++ {
		flag := 0
		for _, j := range nextNode[i] {
			if j < i {
				flag |= 0b1 << res[j]
			}
		}

		for t := 1; t <= 4; t++ {
			if flag&(0b1<<t) == 0 {
				res[i] = t
				break
			}
		}
	}
	return res
}
