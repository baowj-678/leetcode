/**
 * @author: baowj
 * @time: 2022/4/21 19:29:51
 */
package main

import (
	"unicode"
)

func toGoatLatin(sentence string) string {
	ans := ""
	lastIdx := -1
	wordCnt := 1
	for index, c := range sentence {
		if unicode.IsLetter(c) {
			if lastIdx == -1 {
				lastIdx = index
			}
		} else {
			if lastIdx != -1 {
				if wordCnt > 1 {
					ans = ans + " "
				}
				word := sentence[lastIdx:index]
				if isVowel(word[0]) {
					word = word + "ma"
				} else {
					word = word[1:] + string(word[0]) + "ma"
				}
				for i := 0; i < wordCnt; i++ {
					word = word + "a"
				}
				ans += word
				lastIdx = -1
				wordCnt++
			}
		}
	}
	if wordCnt > 1 {
		ans = ans + " "
	}
	word := sentence[lastIdx:]
	if isVowel(word[0]) {
		word = word + "ma"
	} else {
		word = word[1:] + string(word[0]) + "ma"
	}
	for i := 0; i < wordCnt; i++ {
		word = word + "a"
	}
	ans += word
	return ans
}

func isVowel(c uint8) bool {
	vowels := "AEIOUaeiou"
	for i := 0; i < len(vowels); i++ {
		if c == vowels[i] {
			return true
		}
	}
	return false
}

func main() {
	println(toGoatLatin("t"))
}
