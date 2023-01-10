/**
 * @author: baowj
 * @time: 2022/12/18 15:44:31
 * @email: bwj_678@qq.com
 */
package main

func similarPairs(words []string) int {
	res := 0
	wordCntI := 0
	wordCntJ := 0
	for i := 0; i < len(words); i++ {
		wordCntI = 0
		for _, c := range words[i] {
			wordCntI |= 1 << (c - 'a')
		}
		for j := i + 1; j < len(words); j++ {
			wordCntJ = 0
			for _, c := range words[j] {
				wordCntJ |= 1 << (c - 'a')
			}
			if wordCntJ == wordCntI {
				res++
			}
		}
	}
	return res
}
