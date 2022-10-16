/**
 * @author: baowj
 * @time: 2022/10/9 15:59:18
 */
package main

func robotWithString(s string) string {
	flag := make([]byte, len(s))
	lastAlpha := s[len(s)-1]
	for i := len(s) - 1; i >= 0; i-- {
		if s[i] <= lastAlpha {
			flag[i] = s[i]
			lastAlpha = s[i]
		} else {
			flag[i] = lastAlpha
		}
	}
	res := make([]byte, len(s))
	t := []byte{}
	k := 0
	for i := 0; k < len(res) && i < len(flag); i++ {
		for len(t) > 0 && t[len(t)-1] <= flag[i] {
			res[k] = t[len(t)-1]
			k++
			t = t[:len(t)-1]
		}
		if s[i] == flag[i] {
			res[k] = s[i]
			k++
		} else {
			t = append(t, s[i])
		}
	}
	for len(t) > 0 {
		res[k] = t[len(t)-1]
		t = t[:len(t)-1]
		k++
	}
	return string(res)
}
