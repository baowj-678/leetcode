/**
 * @author: baowj
 * @time: 2022/8/6 10:01:32
 */
package main

import (
	"strings"
)

func stringMatching(words []string) []string {
	subStrings := []string{}
	flags := make([]bool, len(words))
	for i := 0; i < len(flags); i++ {
		flags[i] = false
	}
	for i := 0; i < len(words); i++ {
		if flags[i] {
			continue
		}
		for j := i + 1; j < len(words); j++ {
			if flags[j] {
				continue
			}
			if len(words[i]) < len(words[j]) && strings.Contains(words[j], words[i]) {
				subStrings = append(subStrings, words[i])
				flags[i] = true
			} else if strings.Contains(words[i], words[j]) {
				subStrings = append(subStrings, words[j])
				flags[j] = true
			}
		}
	}
	return subStrings
}

func main() {
	ss := []string{"jak", "yjakdn", "tj", "yyjakdn"}
	stringMatching(ss)
}
