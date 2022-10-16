/**
 * @author: baowj
 * @time: 2022/9/27 09:24:11
 */
package main

import (
	"fmt"
	"sort"
)

func CheckPermutation(s1 string, s2 string) bool {
	byte1 := []byte(s1)
	byte2 := []byte(s2)
	if len(s1) != len(s2) {
		return false
	}
	sort.Slice(byte1, func(i, j int) bool {
		return byte1[i] < byte1[j]
	})
	sort.Slice(byte2, func(i, j int) bool {
		return byte2[i] < byte2[j]
	})
	for i := 0; i < len(s1); i++ {
		if byte1[i] != byte2[i] {
			return false
		}
	}
	return true
}

func main() {
	fmt.Println(CheckPermutation("abc", "bca"))
}
