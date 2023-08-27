package main

func circularGameLosers(n int, k int) []int {
	flags := make([]bool, n)
	idx := 0
	i := 1
	for {
		if flags[idx] {
			break
		}
		flags[idx] = true
		idx = (idx + i*k) % n
		i++
	}
	var res []int
	for i, flag := range flags {
		if !flag {
			res = append(res, i+1)
		}
	}
	return res
}
