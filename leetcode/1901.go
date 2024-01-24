package main

import "math"

func findPeakGrid(mat [][]int) []int {
	getMax := func(i int) (int, int) {
		max, idx := math.MinInt, -1
		for k, x := range mat[i] {
			if max < x {
				max, idx = x, k
			}
		}
		return max, idx
	}
	top, bottom := 0, len(mat)-1
	topMax, topIdx := getMax(top)
	bottomMax, bottomIdx := getMax(bottom)
	if (top+1 < len(mat) && topMax > mat[top+1][topIdx]) || (top+1 >= len(mat)) {
		return []int{top, topIdx}
	}
	if (bottom-1 >= 0 && bottomMax > mat[bottom-1][bottomIdx]) || (bottom-1 < 0) {
		return []int{bottom, bottomIdx}
	}

	for top < bottom {
		mid := (top + bottom) / 2
		midMax, midIdx := getMax(mid)
		tFlag := mid+1 < len(mat) && midMax < mat[mid+1][midIdx]
		bFlag := mid-1 >= 0 && midMax < mat[mid-1][midIdx]
		if tFlag {
			top = mid
			topMax, topIdx = midMax, midIdx
		}
		if bFlag {
			bottom = mid
			bottomMax, bottomIdx = midMax, midIdx
		}

		if tFlag && bFlag {
			return []int{mid, midIdx}
		}
	}
	return []int{top, topIdx}
}
