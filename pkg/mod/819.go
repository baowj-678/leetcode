/**
 * @author: baowj
 * @time: 2022/4/17 10:21:06
 */
package main

import (
	"strings"
	"unicode"
)

func mostCommonWord(paragraph string, banned []string) string {
	myMap := make(map[string]int)
	lastIdx := 0
	for i, isContinue := 0, false; i < len(paragraph); i++ {
		if !unicode.IsLetter(rune(paragraph[i])) {
			if isContinue {
				word := strings.ToLower(paragraph[lastIdx:i])
				isContinue = false
				_, ok := myMap[word]
				if ok {
					myMap[word] = myMap[word] + 1
				} else {
					myMap[word] = 1
				}
			}
		} else {
			if !isContinue {
				lastIdx = i
				isContinue = true
			}
		}
	}
	if unicode.IsLetter(rune(paragraph[len(paragraph)-1])) {
		word := strings.ToLower(paragraph[lastIdx:len(paragraph)])
		_, ok := myMap[word]
		if ok {
			myMap[word] = myMap[word] + 1
		} else {
			myMap[word] = 1
		}
	}
	var maxWord string
	max := 0
	for key, value := range myMap {
		isBanned := false
		for _, word := range banned {
			if strings.Compare(key, word) == 0 {
				isBanned = true
				break
			}
		}
		if !isBanned {
			if max < value {
				maxWord = key
				max = value
			}
		}
	}
	return maxWord
}

func main() {
	str := "Bob"
	banned := []string{"hit"}
	println(mostCommonWord(str, banned))
}
