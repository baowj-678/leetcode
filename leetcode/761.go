/**
 * @author: baowj
 * @time: 2022/8/8 10:25:51
 */
package main

import "sort"

func makeLargestSpecial(s string) string {
	if len(s) <= 2 {
		return s
	}
	left, last := 0, 0
	res := ""
	subStrs := []string{}
	for i, c := range s {
		if c == '1' {
			left += 1
		} else {
			left -= 1
		}
		if left == 0 {
			if !(len(subStrs) == 0 && i == len(s)-1) {
				subStrs = append(subStrs, makeLargestSpecial(s[last:i+1]))
				last = i + 1
			}
		}
	}
	if len(subStrs) == 0 {
		res += "1"
		res += makeLargestSpecial(s[1 : len(s)-1])
		res += "0"
	} else {
		sort.Slice(subStrs, func(i, j int) bool {
			return subStrs[i]+subStrs[j] > subStrs[j]+subStrs[i]
		})
		for _, str := range subStrs {
			res += str
		}
	}
	return res
}

func main() {
	println(makeLargestSpecial("11011000"))
}
