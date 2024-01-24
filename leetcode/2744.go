package main

func maximumNumberOfStringPairs(words []string) int {
	cnt := 0
	check := func(a, b string) bool {
		if len(a) != len(b) {
			return false
		}
		for i := 0; i < len(a); i++ {
			if a[i] != b[len(b)-i-1] {
				return false
			}
		}
		return true
	}

	for i := 0; i < len(words); i++ {
		for j := i + 1; j < len(words); j++ {
			if check(words[i], words[j]) {
				cnt++
			}
		}
	}

	return cnt
}
