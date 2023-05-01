/**
 * @author: baowj
 * @time: 2023/4/29 17:09:17
 * @email: bwj_678@qq.com
 */
package main

func longestDecomposition(text string) int {
	checkEqual := func(start, k int) bool {
		end_ := len(text) - 1 - start
		for i := 0; i < k; i++ {
			if text[start+i] != text[end_-k+i+1] {
				return false
			}
		}
		return true
	}

	res := 0
	i := 0
	for i < len(text)/2 {
		ok := false
		for j := i; j < len(text)/2; j++ {
			k := j - i + 1
			if checkEqual(i, k) {
				i = j + 1
				res += 2
				ok = true
				break
			}
		}
		if !ok {
			break
		}
	}
	if res > 0 {
		if i >= (len(text)+1)/2 {
			return res
		} else {
			return res + 1
		}
	}
	return 1
}
