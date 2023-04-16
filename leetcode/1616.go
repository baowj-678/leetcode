/**
 * @author: baowj
 * @time: 2023/3/19 23:12:57
 * @email: bwj_678@qq.com
 */
package main

func checkPalindromeFormation(a string, b string) bool {
	f := func(s1, s2 string) bool {
		for i := 0; i < len(s1)/2; i++ {
			if s1[i] != s2[len(s2)-i-1] {
				isPalindrome := true
				l := i
				r := len(s2) - i - 1
				i, j := l, r
				for i < j {
					if s1[i] != s1[j] {
						isPalindrome = false
						break
					}
					i++
					j--
				}
				if isPalindrome {
					return true
				}
				isPalindrome = true
				i, j = l, r
				for i < j {
					if s2[i] != s2[j] {
						isPalindrome = false
						break
					}
					i++
					j--
				}
				return isPalindrome
			}
		}
		return true
	}

	return f(a, b) || f(b, a)
}
