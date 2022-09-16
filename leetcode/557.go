/**
 * @author: baowj
 * @time: 2022/8/26 21:21:16
 */
package main

func reverseWords(s string) string {
	res := make([]byte, len(s))
	start, end := 0, 0
	for i, c := range s {
		if c == ' ' {
			res[i] = ' '
			end = i - 1
			for start < end {
				res[start] = s[end]
				res[end] = s[start]
				start++
				end--
			}
			start = i + 1
		}
	}
	end = len(s) - 1
	for start < end {
		res[start] = s[end]
		res[end] = s[start]
		start++
		end--
	}
	return string(res)
}
