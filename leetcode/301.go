/**
 * @author: baowj
 * @time: 2023/6/2 16:17
 */
package main

import "strings"

func removeInvalidParentheses(s string) []string {
	l, r := 0, 0
	for _, c := range s {
		if c == '(' {
			l++
		} else if c == ')' {
			if l > 0 {
				l--
			} else {
				r++
			}
		}
	}
	flags := make([]bool, len(s))
	for i := 0; i < len(s); i++ {
		flags[i] = true
	}
	strs := dfs301(s, flags, l, r, 0, 0)
	m := map[string]struct{}{}
	for _, str := range strs {

		m[str] = struct{}{}
	}
	res := make([]string, len(m))
	i := 0
	for k, _ := range m {
		res[i] = k
		i++
	}
	return res
}

func dfs301(s string, flags []bool, l, r int, lp, rp int) []string {
	check := func() bool {
		tmp := 0
		for i, c := range s {
			if flags[i] {
				if c == '(' {
					tmp++
				} else if c == ')' {
					tmp--
				}
				if tmp < 0 {
					return false
				}
			}
		}
		if tmp == 0 {
			return true
		} else {
			return false
		}
	}
	if l+r == 0 && check() {
		builder := strings.Builder{}
		for i, c := range s {
			if flags[i] {
				builder.WriteByte(byte(c))
			}
		}
		return []string{builder.String()}
	}
	res := []string{}
	if l > 0 {
		for i := lp; i < len(s); i++ {
			if s[i] == '(' && flags[i] == true {
				flags[i] = false
				tmp := dfs301(s, flags, l-1, r, i+1, rp)
				res = append(res, tmp...)
				flags[i] = true
			}
		}
	}

	if r > 0 {
		for i := rp; i < len(s); i++ {
			if s[i] == ')' && flags[i] == true {
				flags[i] = false
				tmp := dfs301(s, flags, l, r-1, lp, i+1)
				res = append(res, tmp...)
				flags[i] = true
			}
		}
	}
	return res
}
