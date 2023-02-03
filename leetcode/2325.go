/**
 * @author: baowj
 * @time: 2023/2/3 09:40:10
 * @email: bwj_678@qq.com
 */
package main

func decodeMessage(key string, message string) string {
	m := [26]byte{}
	alpha := 0
	for _, c := range key {
		if c == ' ' {
			continue
		}
		if m[c-'a'] == 0 {
			m[c-'a'] = byte(alpha + 'a')
			alpha++
		}
	}

	res := []byte(message)
	for i, c := range res {
		if c != ' ' {
			res[i] = m[c-'a']
		}
	}

	return string(res)
}
