package main

func maximumSumOfHeightsI(maxHeights []int) int64 {
	left := make([]int64, len(maxHeights))
	for i := 0; i < len(maxHeights); i++ {
		v, m := int64(0), maxHeights[i]
		for j := i; j >= 0; j-- {
			if m >= maxHeights[j] {
				v += int64(maxHeights[j])
				m = maxHeights[j]
			} else {
				v += int64(m)
			}
		}
		left[i] = v
	}
	right := make([]int64, len(maxHeights))
	for i := 0; i < len(maxHeights); i++ {
		v, m := int64(0), maxHeights[i]
		for j := i; j < len(maxHeights); j++ {
			if m >= maxHeights[j] {
				v += int64(maxHeights[j])
				m = maxHeights[j]
			} else {
				v += int64(m)
			}
		}
		right[i] = v
	}

	res := int64(0)
	for i := 0; i < len(maxHeights); i++ {
		if left[i]+right[i]-int64(maxHeights[i]) > res {
			res = left[i] + right[i] - int64(maxHeights[i])
		}
	}

	return res
}
