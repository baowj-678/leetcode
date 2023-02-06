/**
 * @author: baowj
 * @time: 2023/2/6 22:03:11
 * @email: bwj_678@qq.com
 */
package main

func vowelStrings(words []string, queries [][]int) []int {
	vowel := [26]bool{}
	vowel[int('a'-'a')] = true
	vowel[int('e'-'a')] = true
	vowel[int('i'-'a')] = true
	vowel[int('o'-'a')] = true
	vowel[int('u'-'a')] = true
	prefixSum := make([]int, len(words))
	isVowelWord := func(s string) bool {
		return vowel[s[0]-'a'] && vowel[s[len(s)-1]-'a']
	}
	tmpSum := 0
	for i, word := range words {
		if isVowelWord(word) {
			tmpSum++
		}
		prefixSum[i] = tmpSum
	}

	res := make([]int, len(queries))
	for i, query := range queries {
		if query[0] == 0 {
			res[i] = prefixSum[query[1]]
		} else {
			res[i] = prefixSum[query[1]] - prefixSum[query[0]]
		}
	}
	return res
}
