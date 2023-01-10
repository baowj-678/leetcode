/**
 * @author: baowj
 * @time: 2022/12/5 22:46:45
 * @email: bwj_678@qq.com
 */
package main

import "strings"

func isCircularSentence(sentence string) bool {
	words := strings.Split(sentence, " ")
	for i := 0; i < len(words); i++ {
		word1 := words[(i-1+len(words))%len(words)]
		word2 := words[i]
		if word2[0] != word1[len(word1)-1] {
			return false
		}
	}
	return true
}

func main() {
	println(-1 % 6)
}
