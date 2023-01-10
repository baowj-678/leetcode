/**
 * @author: baowj
 * @time: 2022/10/23 11:51:51
 */
package main

func mergeAlternately(word1 string, word2 string) string {
	res := make([]byte, len(word1)+len(word2))
	i, j, k := 0, 0, 0
	for i < len(word1) && j < len(word2) {
		res[k] = word1[i]
		i++
		k++
		res[k] = word2[j]
		j++
		k++
	}
	if i == len(word1) {
		for j < len(word2) {
			res[k] = word2[j]
			k++
			j++
		}
	} else if j == len(word2) {
		for i < len(word1) {
			res[k] = word1[i]
			k++
			i++
		}
	}
	return string(res)
}
