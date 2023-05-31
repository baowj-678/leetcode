/**
 * @author: baowj
 * @time: 2023/5/25 10:24
 */
package main

func oddString(words []string) string {
	difference := make([][]int, len(words))
	for i, word := range words {
		difference[i] = make([]int, len(word))
		for j := 1; j < len(word); j++ {
			difference[i][j-1] = int(word[j] - word[j-1])
		}
	}
	check := func(a, b []int) bool {
		for i := 0; i < len(a); i++ {
			if a[i] != b[i] {
				return false
			}
		}
		return true
	}
	for i := 0; i < len(words); i++ {
		cnt := 0
		for j := 0; j < len(words); j++ {
			if !check(difference[i], difference[j]) {
				cnt++
			}
			if cnt > 1 {
				return words[i]
			}
		}
	}
	return words[0]
}
