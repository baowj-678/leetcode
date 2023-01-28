/**
 * @author: baowj
 * @time: 2023/1/28 11:17:11
 * @email: bwj_678@qq.com
 */
package main

func greatestLetter(s string) string {
	littleLetter := [26]bool{}
	bigLetter := [26]bool{}

	for _, c := range s {
		if 'a' <= c && c <= 'z' {
			littleLetter[c-'a'] = true
		} else {
			bigLetter[c-'A'] = true
		}
	}

	for i := 25; i >= 0; i-- {
		if littleLetter[i] && bigLetter[i] {
			return string(byte(i) + 'A')
		}
	}
	return ""
}
