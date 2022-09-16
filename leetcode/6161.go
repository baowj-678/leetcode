/**
 * @author: baowj
 * @time: 2022/8/28 21:34:59
 */
package main

func removeStars(s string) string {
	res := []byte{}
	for i, j := len(s)-1, len(s)-1; i >= 0; {
		if s[j] == '*' {
			for i >= 0 && s[i] == '*' {
				i--
			}
			i--
			j--
		} else {
			if j > i {
				j--
			} else {
				res = append(res, s[i])
				j--
				i--
			}
		}
	}
	for i := 0; i < len(res)/2; i++ {
		res[i], res[len(res)-i-1] = res[len(res)-i-1], res[i]
	}
	return string(res)
}
