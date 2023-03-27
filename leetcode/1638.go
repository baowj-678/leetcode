/**
 * @author: baowj
 * @time: 2023/3/27 13:56
 */
package main

func countSubstrings(s string, t string) int {
	check := func(a, b string) bool {
		cnt := 0
		for i := 0; i < len(a); i++ {
			if a[i] != b[i] {
				cnt++
			}
		}
		return cnt == 1
	}
	res := 0
	for l := 1; l <= len(s); l++ {
		for left := 0; left+l <= len(s); left++ {
			for tLeft := 0; tLeft+l <= len(t); tLeft++ {
				if check(s[left:left+l], t[tLeft:tLeft+l]) {
					res++
				}
			}
		}
	}
	return res
}
