/**
 * @author: baowj
 * @time: 2023/3/19 20:26:14
 * @email: bwj_678@qq.com
 */
package main

func findLexSmallestString(s string, a int, b int) string {
	flags := make([]bool, len(s))

	minusB := func(s string, isOdd bool) string {
		res := []byte(s)
		// odd
		min := s[1] - '0'
		v := byte(0)
		for i := 0; i < 10; i++ {
			tmp := (s[1] - '0' + byte(i)*byte(a)) % 10
			if tmp < min {
				min = tmp
				v = (byte(i) * byte(a)) % 10
			}
		}
		for i := 1; i < len(s); i += 2 {
			res[i] = '0' + (res[i]-'0'+v)%10
		}
		// even
		if isOdd {
			min = s[0] - '0'
			v = byte(0)
			for i := 0; i < 10; i++ {
				tmp := (s[0] - '0' + byte(i)*byte(a)) % 10
				if tmp < min {
					min = tmp
					v = (byte(i) * byte(a)) % 10
				}
			}
			for i := 0; i < len(s); i += 2 {
				res[i] = '0' + (res[i]-'0'+v)%10
			}
		}
		return string(res)
	}

	isOdd := b%2 == 1
	res := s
	for i := 0; i < len(s); i++ {
		p := (i * b) % len(s)
		if flags[p] {
			break
		}
		flags[p] = true
		tmps := s[p:] + s[:p]
		tmps = minusB(tmps, isOdd)
		if tmps < res {
			res = tmps
		}
	}
	return res
}
