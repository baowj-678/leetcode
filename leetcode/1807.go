/**
 * @author: baowj
 * @time: 2023/1/12 17:20:35
 * @email: bwj_678@qq.com
 */
package main

func evaluate(s string, knowledge [][]string) string {
	m := map[string]string{}
	for _, k := range knowledge {
		m[k[0]] = k[1]
	}
	last, left := 0, -1
	res := ""
	for i, c := range s {
		if c == '(' {
			res += s[last:i]
			left = i + 1
			continue
		} else if c == ')' {
			key := s[left:i]
			if v, ok := m[key]; ok {
				res += v
			} else {
				res += "?"
			}
			last = i + 1
			left = -1
		}
	}
	if last != len(s) {
		res += s[last:]
	}
	return res
}
