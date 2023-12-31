package main

func maxStudents(seats [][]byte) int {
	l := len(seats[0])
	dp := make([]int, 1<<l)
	for _, seat := range seats {
		s := byte(0)
		for _, tmp := range seat {
			switch tmp {
			case '.':
				s = (s << 1) | 1
			case '#':
				s = s << 1
			}
		}
		newDp := make([]int, 1<<l)

		for iInt := 0; iInt < 1<<l; iInt++ {
			i := byte(iInt)
			if (i^s)&i != 0 || i&(i<<1) != 0 || i&(i>>1) != 0 {
				continue
			}
			add := 0
			for k := 0; k < l; k++ {
				if (1<<k)&i != 0 {
					add += 1
				}
			}

			for jInt := 0; jInt < 1<<l; jInt++ {
				j := byte(jInt)
				if i&(j<<1) != 0 || i&(j>>1) != 0 {
					continue
				}

				if dp[j]+add > newDp[i] {
					newDp[i] = dp[j] + add
				}
			}
		}
		dp = newDp
	}
	res := 0
	for _, x := range dp {
		if x > res {
			res = x
		}
	}
	return res
}
