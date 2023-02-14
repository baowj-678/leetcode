package main

func areAlmostEqual(s1 string, s2 string) bool {
	first, second := -1, -1
	if len(s1) != len(s2) {
		return false
	}
	for i := 0; i < len(s1); i++ {
		if s1[i] != s2[i] {
			if first == -1 {
				first = i
			} else if second == -1 {
				second = i
			} else {
				return false
			}
		}
	}
	if first != -1 && second != -1 {
		if s1[second] == s2[first] && s1[first] == s2[second] {
			return true
		}
	} else if first == -1 && second == -1 {
		return true
	}
	return false
}
