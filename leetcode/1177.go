/**
 * @author: baowj
 * @time: 2023/6/15 15:48
 */
package main

import "fmt"

func canMakePaliQueries(s string, queries [][]int) []bool {
	prefix := make([][26]int, len(s))
	prefix[0][s[0]-'a']++
	for i := 1; i < len(s); i++ {
		prefix[i][s[i]-'a']++
		for k := 0; k < 26; k++ {
			prefix[i][k] += prefix[i-1][k]
		}
	}

	check := func(l, r, k int) bool {
		length := r - l + 1
		alphabet := [26]int{}
		for i := 0; i < 26; i++ {
			alphabet[i] = prefix[r][i]
		}
		if l > 0 {
			for i := 0; i < 26; i++ {
				alphabet[i] -= prefix[l-1][i]
			}
		}
		for i := 0; i < 26; i++ {
			length -= (alphabet[i] / 2) * 2
		}
		return length/2 <= k
	}
	res := make([]bool, len(queries))
	for i, q := range queries {
		res[i] = check(q[0], q[1], q[2])
	}
	return res
}

func main() {
	s := "abcda"
	queries := [][]int{{3, 3, 0}, {1, 2, 0}, {0, 3, 1}, {0, 3, 2}, {0, 4, 1}}
	fmt.Println(canMakePaliQueries(s, queries))
}
