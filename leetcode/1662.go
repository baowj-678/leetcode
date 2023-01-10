/**
 * @author: baowj
 * @time: 2022/11/1 16:26:30
 */
package main

func arrayStringsAreEqual(word1 []string, word2 []string) bool {
	str1 := ""
	for _, word := range word1 {
		str1 += word
	}
	str2 := ""
	for _, word := range word2 {
		str2 += word
	}
	return str1 == str2
}
