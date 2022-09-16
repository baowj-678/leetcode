/**
 * @author: baowj
 * @time: 2022/9/8 09:26:32
 */
package main

func constructArray(n int, k int) []int {
	res := make([]int, n)
	if k%2 == 0 {
		j := (k+1)/2 + 1
		step := 0
		for i := 0; i < k+1; i++ {
			j += step
			res[i] = j
			if step < 0 {
				step = -step + 1
			} else {
				step = -step - 1
			}
		}
	} else {
		j := (k + 1) / 2
		step := 0
		for i := 0; i < k+1; i++ {
			j += step
			res[i] = j
			if step <= 0 {
				step = -step + 1
			} else {
				step = -step - 1
			}
		}
	}

	for i := k + 1; i < n; i++ {
		res[i] = i + 1
	}
	return res
}
