package main

import "strings"

func repeatLimitedString(s string, repeatLimit int) string {
	charNums := make([]int, 26)
	for _, c := range s {
		charNums[c-'a']++
	}
	builder := strings.Builder{}
	lastChar, lastLen := -1, 0
	for {
		flag := false
		for i := len(charNums) - 1; i >= 0; i-- {
			if charNums[i] > 0 && (lastChar != i || lastLen < repeatLimit) {
				builder.WriteByte(byte('a' + i))
				charNums[i]--

				if lastChar == i {
					lastLen++
				} else {
					lastLen = 1
				}
				lastChar = i
				flag = true
				break
			}
		}

		if !flag {
			break
		}

		for len(charNums) > 0 && charNums[len(charNums)-1] == 0 {
			charNums = charNums[:len(charNums)-1]
		}
	}

	return builder.String()
}

func main() {
	repeatLimitedString("cczazcc", 3)
}
