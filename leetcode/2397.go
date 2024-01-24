package main

func maximumRows(matrix [][]int, numSelect int) int {
	n, m := len(matrix), len(matrix[0])
	matrixNum := make([]int, n)
	for i := 0; i < n; i++ {
		x := 0
		for j, num := range matrix[i] {
			if num == 1 {
				x ^= 1 << j
			}
		}
		matrixNum[i] = x
	}
	res := 0
	for i := 0; i < 1<<m; i++ {
		cnt := 0
		for j := 0; j < m; j++ {
			if i&(1<<j) == 0 {
				cnt++
			}
		}
		if cnt != numSelect {
			continue
		}
		cnt = 0
		for _, x := range matrixNum {
			if x&i == 0 {
				cnt++
			}
		}
		if cnt > res {
			res = cnt
		}
	}
	return res
}
