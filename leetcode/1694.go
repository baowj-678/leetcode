/**
 * @author: baowj
 * @time: 2022/10/1 09:07:09
 */
package main

func reformatNumber(number string) string {
	cnt := 0
	for _, c := range number {
		if c != ' ' && c != '-' {
			cnt++
		}
	}
	tmp := cnt
	for tmp > 4 {
		tmp -= 3
		cnt++
	}
	hasTwo := false
	if tmp == 4 {
		cnt++
		hasTwo = true
	}
	res := make([]byte, cnt)
	cnt = 0
	i := 0
	for _, c := range number {
		if c != ' ' && c != '-' {
			res[i] = byte(c)
			i++
			cnt++
			if hasTwo && i == len(res)-3 {
				res[i] = '-'
				cnt = 0
				i++
			}
			if cnt >= 3 && i < len(res) {
				res[i] = '-'
				i++
				cnt = 0
			}
		}
	}
	return string(res)
}
