/**
 * @author: baowj
 * @time: 2023/2/25 10:48:58
 * @email: bwj_678@qq.com
 */
package main

func minimumSwap(s1 string, s2 string) int {
	if len(s1) != len(s2) {
		return -1
	}
	m1, m2 := 0, 0
	for i := 0; i < len(s1); i++ {
		if s1[i] == s2[i] {
			continue
		}
		if s1[i] == 'x' {
			m1++
		} else {
			m2++
		}
	}
	if (m1-m2)%2 != 0 {
		return -1
	}
	res := m1/2 + m2/2
	if m1%2 == 1 {
		res += 2
	}
	return res
}
