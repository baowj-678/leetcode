package main

func maxPalindromes(s string, k int) int {
	last := -1
	cnt := 0
	for i := 0; i < len(s); i++ {
		if i <= last {
			continue
		}
		if i+k <= len(s) {
			isOk := true
			for l, r := i, i+k-1; l < r; {
				if s[l] != s[r] {
					isOk = false
					break
				}
				l++
				r--
			}
			if isOk {
				cnt++
				last = i + k - 1
				continue
			}
		}
		if i+k+1 <= len(s) {
			isOk := true
			for l, r := i, i+k; l < r; {
				if s[l] != s[r] {
					isOk = false
					break
				}
				l++
				r--
			}
			if isOk {
				cnt++
				last = i + k
				continue
			}
		}
	}
	return cnt
}
