/**
 * @author: baowj
 * @time: 2022/8/21 10:49:31
 */
package main

func shiftingLetters(s string, shifts [][]int) string {
	ss := make([]int32, len(s)+5)
	res := make([]byte, len(s))
	for _, shift := range shifts {
		if shift[2] == 1 {
			ss[shift[0]] = (ss[shift[0]] + 1) % 26
			ss[shift[1]+1] = (ss[shift[1]+1] - 1) % 26
		} else {
			ss[shift[0]] = (ss[shift[0]] - 1) % 26
			ss[shift[1]+1] = (ss[shift[1]+1] + 1) % 26
		}
	}
	for i := 1; i < len(s); i++ {
		ss[i] = (ss[i-1] + ss[i]) % 26
	}
	for i, c := range s {
		res[i] = byte('a' + (c-'a'+ss[i])%26)
	}
	return string(res)
}

func main() {
	println(-1 % 26)
}
