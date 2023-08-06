package main

func countDifferentSubsequenceGCDs(nums []int) int {
	res := 0
	table := make([]bool, 2*1e5+1)
	max := 0
	for _, num := range nums {
		table[num] = true
		if num > max {
			max = num
		}
	}

	gcd := func(x, y int) int {
		if x < y {
			x, y = y, x
		}
		for y != 0 {
			x, y = y, x%y
		}
		return x
	}

	check := func(x int) bool {
		v := 0
		for i := 1; i*x <= max; i++ {
			if table[i*x] {
				if v != 0 {
					v = gcd(v, i*x)
				} else {
					v = i * x
				}
			}
		}
		if v == x {
			return true
		} else {
			return false
		}
	}

	for i := 1; i <= max; i++ {
		if check(i) {
			res++
		}
	}
	return res
}
