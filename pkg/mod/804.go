/**
 * @author: baowj
 * @time: 2022/4/10 10:41:47
 */
package main

import (
	"sort"
)

type IntSlice []int64

func (s IntSlice) Len() int           { return len(s) }
func (s IntSlice) Swap(i, j int)      { s[i], s[j] = s[j], s[i] }
func (s IntSlice) Less(i, j int) bool { return s[i] < s[j] }

func uniqueMorseRepresentations(words []string) int {
	table := []string{".-", "-...", "-.-.", "-..", ".", "..-.", "--.", "....", "..", ".---", "-.-", ".-..", "--", "-.", "---", ".--.", "--.-", ".-.", "...", "-", "..-", "...-", ".--", "-..-", "-.--", "--.."}

	codes := make([]int64, len(words))
	for index, word := range words {
		i := 0
		var val int64 = 0
		for _, alpha := range word {
			alpha -= 'a'
			for _, code := range table[alpha] {
				if code == '-' {
					val |= 1 << i
				}
				i++
			}
		}
		codes[index] = val
	}
	sort.Sort(IntSlice(codes))
	cnt := 1
	last := codes[0]
	for _, code := range codes {
		if code != last {
			cnt++
			last = code
		}
	}
	return cnt
}

func main() {
	words := []string{"zocd", "gjkl", "hzqk", "hzgq", "gjkl"}
	println(uniqueMorseRepresentations(words))
}
