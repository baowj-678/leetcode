/**
 * @author: baowj
 * @time: 2023/1/17 21:10:24
 * @email: bwj_678@qq.com
 */
package main

import "strings"

func areSentencesSimilar(sentence1 string, sentence2 string) bool {
	word1 := strings.Split(sentence1, " ")
	word2 := strings.Split(sentence2, " ")
	l := 0
	for ; l < len(word2) && l < len(word1); l++ {
		if word1[l] != word2[l] {
			break
		}
	}
	r1, r2 := len(word1)-1, len(word2)-1
	for r1 >= 0 && r2 >= 0 {
		if word1[r1] != word2[r2] {
			break
		}
		r1--
		r2--
	}
	if l > r1 || l > r2 {
		return true
	}
	return false
}
