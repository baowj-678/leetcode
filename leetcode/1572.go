package main

func diagonalSum(mat [][]int) int {
	res := 0
	for i := 0; i < len(mat); i++ {
		res += mat[i][i]
		res += mat[i][len(mat)-1-i]
	}
	if len(mat)%2 == 1 {
		res -= mat[len(mat)/2][len(mat)/2]
	}
	return res
}
