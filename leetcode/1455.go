/**
 * @author: baowj
 * @time: 2022/8/21 10:16:21
 */
package main

func isPrefixOfWord(sentence string, searchWord string) int {
	idx := 1
	searchIdx := 0
	for _, c := range sentence {
		if c == ' ' {
			idx++
			searchIdx = 0
		} else {
			if searchIdx != -1 {
				if searchWord[searchIdx] == uint8(c) {
					searchIdx++
				} else {
					searchIdx = -1
				}
				if searchIdx == len(searchWord) {
					return idx
				}
			}
		}
	}
	return -1
}
