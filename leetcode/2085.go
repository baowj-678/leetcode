package main

func countWords(words1 []string, words2 []string) int {
	m := map[string]int{}
	for _, word := range words1 {
		if v, ok := m[word]; ok {
			m[word] = v + 1
		} else {
			m[word] = 1
		}
	}

	res := map[string]int{}
	for _, word := range words2 {
		if v, ok := m[word]; ok && v == 1 {
			if v, ok = res[word]; ok {
				res[word] = v + 1
			} else {
				res[word] = 1
			}
		}
	}

	cnt := 0
	for _, v := range res {
		if v == 1 {
			cnt++
		}
	}
	return cnt
}
