/**
 * @author: baowj
 * @time: 2022/10/9 15:31:18
 */
package main

func scoreOfParentheses(s string) int {
	if s == "()" {
		return 1
	}
	l, lastIdx, sum := 0, 0, 0
	for i := 0; i < len(s); i++ {
		if s[i] == '(' {
			l++
		} else {
			l--
		}
		if l == 0 {
			if lastIdx+1 < i {
				sum += 2 * scoreOfParentheses(s[lastIdx+1:i])
			} else {
				sum += scoreOfParentheses(s[lastIdx : i+1])
			}
			lastIdx = i + 1
		}
	}
	return sum
}
