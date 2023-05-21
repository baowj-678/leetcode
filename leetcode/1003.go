/**
 * @author: baowj
 * @time: 2023/5/3 10:40:32
 * @email: bwj_678@qq.com
 */
package main

func isValid(s string) bool {
	stack := []byte{}
	check := func(s []byte) bool {
		if len(s) < 3 {
			return false
		}
		if s[len(s)-1] == 'c' && s[len(s)-2] == 'b' && s[len(s)-3] == 'a' {
			return true
		}
		return false
	}
	for _, c := range s {
		stack = append(stack, byte(c))
		if check(stack) {
			stack = stack[:len(stack)-3]
		}
	}

	return len(stack) == 0
}
