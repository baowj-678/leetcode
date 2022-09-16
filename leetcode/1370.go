/**
 * @author: baowj
 * @time: 2022/9/5 11:04:20
 */
package main

func sortString(s string) string {
	table := [26]int{}
	for _, c := range s {
		table[c-'a']++
	}
	res := make([]byte, len(s))
	p := 0
	for p < len(s) {
		for i := 0; i < 26; i++ {
			if table[i] != 0 {
				table[i]--
				res[p] = byte('a' + i)
				p++
			}
		}

		for i := 25; i >= 0; i-- {
			if table[i] != 0 {
				table[i]--
				res[p] = byte('a' + i)
				p++
			}
		}
	}
	return string(res)
}
