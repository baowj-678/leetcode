/**
 * @author: baowj
 * @time: 2023/2/27 14:08
 */
package main

func maxScoreWords(words []string, letters []byte, score []int) int {
	maxScore := 0
	letterCnt := [26]byte{}
	for _, letter := range letters {
		letterCnt[letter-'a']++
	}

	letterCmp := func(a, b [26]byte) bool {
		for i := 0; i < len(a); i++ {
			if a[i] > b[i] {
				return false
			}
		}
		return true
	}
	letterAdd := func(a, b [26]byte) [26]byte {
		for i := 0; i < 26; i++ {
			a[i] += b[i]
		}
		return a
	}
	letterScore := func(a [26]byte) int {
		res := 0
		for i := 0; i < len(a); i++ {
			res += int(a[i]) * score[i]
		}
		return res
	}

	newWords := []string{}
	wordLetters := [][26]byte{}
	for _, word := range words {
		aWordLetters := [26]byte{}
		for _, c := range word {
			aWordLetters[c-'a']++
		}
		if letterCmp(aWordLetters, letterCnt) {
			newWords = append(newWords, word)
			wordLetters = append(wordLetters, aWordLetters)
		}
	}

	for i := 1; i < (1<<len(newWords))-1; i++ {
		tmpLetter := [26]byte{}
		for k := 0; k < len(newWords); k++ {
			if i&(1<<k) > 0 {
				tmpLetter = letterAdd(tmpLetter, wordLetters[k])
			}
		}
		if letterCmp(tmpLetter, letterCnt) {
			s := letterScore(tmpLetter)
			if s > maxScore {
				maxScore = s
			}
		}
	}
	return maxScore
}

func main() {
	words := []string{"baaa", "aacc", "ccbc", "da", "dbbc"}
	letters := []byte{'a', 'a', 'a', 'c', 'c', 'c', 'c', 'c', 'd', 'd', 'd'}
	score := []int{9, 4, 10, 8, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}
	print(maxScoreWords(words, letters, score))
}
