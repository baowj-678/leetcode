package main

func halvesAreAlike(s string) bool {
	count := 0
	for i := 0; i < len(s)/2; i++ {
		tmp := s[i]
		if tmp < 'a' {
			tmp += 'a' - 'A'
		}

		if tmp == 'a' || tmp == 'e' || tmp == 'i' || tmp == 'o' || tmp == 'u' {
			count++
		}
	}
	for i := len(s) / 2; i < len(s); i++ {
		tmp := s[i]
		if tmp < 'a' {
			tmp += 'a' - 'A'
		}

		if tmp == 'a' || tmp == 'e' || tmp == 'i' || tmp == 'o' || tmp == 'u' {
			count--
		}
	}
	if count == 0 {
		return true
	}
	return false
}
