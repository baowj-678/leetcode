package main

import (
	"fmt"
	"strconv"
)

func ambiguousCoordinates(s string) []string {
	s = s[1 : len(s)-1]
	res := []string{}
	for i := 1; i < len(s); i++ {
		res = subfunc(s[:i], s[i:len(s)], res)
	}
	return res
}

func subfunc(str1, str2 string, strs []string) []string {
	strSet1 := []string{}
	if check816(str1) {
		strSet1 = append(strSet1, str1)
	}
	for i := 1; i < len(str1); i++ {
		s := str1[:i] + "." + str1[i:]
		if check816(s) {
			strSet1 = append(strSet1, s)
		}
	}

	if len(strSet1) == 0 {
		return strs
	}

	if check816(str2) {
		for i := 0; i < len(strSet1); i++ {
			strs = append(strs, "("+strSet1[i]+", "+str2+")")
		}
	}
	for i := 1; i < len(str2); i++ {
		s := str2[:i] + "." + str2[i:]
		if check816(s) {
			for j := 0; j < len(strSet1); j++ {
				strs = append(strs, "("+strSet1[j]+", "+s+")")
			}
		}
	}
	return strs
}

func check816(s string) bool {
	prefix := 0
	for i := 0; i < len(s); i++ {
		if s[i] == '0' {
			prefix++
			if i == len(s)-1 && prefix > 1 {
				return false
			}
		} else if s[i] == '.' {
			if prefix > 1 {
				return false
			} else {
				break
			}
		} else {
			if prefix == 0 {
				break
			} else {
				return false
			}
		}
	}
	prefix = len(s) - 1
	for i := len(s) - 1; i >= 0; i-- {
		if s[i] == '.' {
			prefix = i
		}
	}
	subfix := 0
	for i := len(s) - 1; i >= prefix; i-- {
		if s[i] == '0' {
			subfix++
		} else {
			if subfix > 0 {
				return false
			} else {
				break
			}
		}
	}
	return true
}

func main() {
	fmt.Println(strconv.Atoi("00.0"))
}
