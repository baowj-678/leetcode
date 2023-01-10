/**
 * @author: baowj
 * @time: 2022/11/1 16:47:54
 */
package main

func letterCasePermutation(s string) []string {
	res := []string{}
	res = dfs784(res, 0, []byte(s))
	return res
}

func dfs784(res []string, idx int, s []byte) []string {
	if idx >= len(s) {
		res = append(res, string(s))
		return res
	}
	if s[idx] >= 'A' {
		res = dfs784(res, idx+1, s)
		if s[idx] >= 'a' {
			s[idx] -= 'a' - 'A'
		} else {
			s[idx] += 'a' - 'A'
		}
		res = dfs784(res, idx+1, s)
	} else {
		res = dfs784(res, idx+1, s)
	}
	return res
}
