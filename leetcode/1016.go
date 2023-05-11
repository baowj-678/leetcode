/**
 * @author: baowj
 * @time: 2023/5/11 10:07
 */
package main

import "strconv"

func queryString(s string, n int) bool {
	if n == 1 {
		for i := 0; i < len(s); i++ {
			if s[i] == '1' {
				return true
			}
		}
		return false
	}
	k := 0
	tmp := n
	for tmp > 0 {
		tmp >>= 1
		k++
	}
	m := map[int]struct{}{}
	left := 1 << (k - 1)
	for i := 0; i+k <= len(s); i++ {
		sub := s[i : i+k]
		x, err := strconv.ParseInt(sub, 2, 0)
		if err != nil {
			return false
		}
		if int(x) >= left && int(x) <= n {
			m[int(x)] = struct{}{}
		}
	}
	if len(m) != n-left+1 {
		return false
	}

	k--
	left = 1 << (k - 1)
	right := (1 << k) - 1

	m = map[int]struct{}{}
	for i := 0; i+k <= len(s); i++ {
		sub := s[i : i+k]
		x, err := strconv.ParseInt(sub, 2, 0)
		if err != nil {
			return false
		}
		if int(x) >= left && int(x) <= right {
			m[int(x)] = struct{}{}
		}
	}
	if len(m) != right-left+1 {
		return false
	}
	return true
}
