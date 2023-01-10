package main

func countConsistentStrings(allowed string, words []string) int {
	flag := make([]int, 26)
	for _, x := range allowed {
		flag[x-'a']++
	}
	res := 0
	for _, word := range words {
		isOK := true
		for _, x := range word {
			if flag[x-'a'] == 0 {
				isOK = false
			}
		}
		if isOK {
			res++
		}
	}
	return res
}
