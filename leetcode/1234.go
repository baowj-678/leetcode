/**
 * @author: baowj
 * @time: 2023/2/13 12:45
 */
package main

func balancedString(s string) int {
	leftcnt, rightcnt := make([][4]int, len(s)), make([][4]int, len(s))
	alpha := []int32{'Q', 'W', 'E', 'R'}
	for i, c := range s {
		if i > 0 {
			for j, a := range alpha {
				if c == a {
					leftcnt[i][j] = leftcnt[i-1][j] + 1
				} else {
					leftcnt[i][j] = leftcnt[i-1][j]
				}
			}
		} else {
			for j, a := range alpha {
				if c == a {
					leftcnt[i][j] = 1
				} else {
					leftcnt[i][j] = 0
				}
			}
		}
	}

	for i := len(s) - 1; i >= 0; i-- {
		c := int32(s[i])
		if i < len(s)-1 {
			for j, a := range alpha {
				if c == a {
					rightcnt[i][j] = rightcnt[i+1][j] + 1
				} else {
					rightcnt[i][j] = rightcnt[i+1][j]
				}
			}
		} else {
			for j, a := range alpha {
				if c == a {
					rightcnt[i][j] = 1
				} else {
					rightcnt[i][j] = 0
				}
			}
		}
	}
	check := func(nums [4]int) bool {
		for _, n := range nums {
			if n > len(s)/4 {
				return false
			}
		}
		return true
	}
	l, r := -1, 0
	res := len(s)
	for l < len(s) && r < len(s) && l < r {
		sum := [4]int{}
		if l >= 0 {
			sum = leftcnt[l]
		}
		for i := 0; i < 4; i++ {
			sum[i] += rightcnt[r][i]
		}
		if check(sum) {
			if r-l-1 < res {
				res = r - l - 1
			}
			l++
		} else {
			r++
		}
	}
	for l < len(s) {
		if check(leftcnt[l]) {
			if len(s)-l-1 < res {
				res = len(s) - l - 1
			}
		} else {
			break
		}
		l++
	}
	return res
}
