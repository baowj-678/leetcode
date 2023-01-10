/**
 * @author: baowj
 * @time: 2023/1/8 11:38:21
 * @email: bwj_678@qq.com
 */
package main

func prefixCount(words []string, pref string) int {
	cnt := 0
	for _, word := range words {
		if contains2185(word, pref) {
			cnt++
		}
	}
	return cnt
}

func contains2185(word string, pref string) bool {
	if len(word) < len(pref) {
		return false
	}

	for i := 0; i < len(pref); i++ {
		if word[i] != pref[i] {
			return false
		}
	}
	return true
}
