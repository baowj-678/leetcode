package main

func addMinimum(word string) int {
	should := 'a'
	res := 0
	for _, c := range word {
		if c == should {
			should = 'a' + (c+1-'a')%3
			continue
		}
		if c > should {
			res += int(c - should)
			should = 'a' + (c+1-'a')%3
			continue
		}
		if c < should {
			res += int('c' - should + 1)
			res += int(c - 'a')
			should = 'a' + (c+1-'a')%3
		}
	}
	if should != 'a' {
		res += int('c' - should + 1)
	}
	return res
}
