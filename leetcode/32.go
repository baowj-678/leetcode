/**
 * @author: baowj
 * @time: 2022/8/14 10:08:58
 */
package main

func longestValidParentheses(s string) int {
	maxLen, lastIdx, count := 0, 0, 0
	for i := 0; i < len(s); i++ {
		if s[i] == '(' {
			count += 1
		} else {
			count -= 1
		}
		if count == 0 && i > lastIdx {
			tmpLen := i - lastIdx + 1
			if tmpLen > maxLen {
				maxLen = tmpLen
			}
		} else if count < 0 {
			lastIdx = i + 1
			count = 0
		}
	}

	count, lastIdx = 0, len(s)-1
	for i := len(s) - 1; i >= 0; i-- {
		if s[i] == ')' {
			count += 1
		} else {
			count -= 1
		}
		if count == 0 && i < lastIdx {
			tmpLen := lastIdx - i + 1
			if tmpLen > maxLen {
				maxLen = tmpLen
			}
		} else if count < 0 {
			lastIdx = i - 1
			count = 0
		}
	}
	return maxLen
}

func main() {
	print(longestValidParentheses(
		")()())"))
}
