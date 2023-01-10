/**
 * @author: baowj
 * @time: 2022/12/10 17:30:15
 * @email: bwj_678@qq.com
 */
package main

import "strconv"

func numDifferentIntegers(word string) int {
	word += "."
	m := map[string]struct{}{}
	last := -1
	for i, c := range word {
		if c >= '0' && c <= '9' {
			if last == -1 {
				last = i
			}
		} else {
			if last != -1 {
				m[f1805(word[last:i])] = struct{}{}
			}
			last = -1
		}
	}
	return len(m)
}

func f1805(s string) string {
	for i, c := range s {
		if c != '0' {
			return s[i:len(s)]
		}
	}
	return "0"
}

func main() {
	println(strconv.Atoi("000"))
}
