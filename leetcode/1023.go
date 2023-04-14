/**
 * @author: baowj
 * @time: 2023/4/14 10:09
 */
package main

import "unicode"

func camelMatch(queries []string, pattern string) []bool {
	match := func(query string) bool {
		pi := 0
		for _, c := range query {
			if unicode.IsLower(c) {
				if pi < len(pattern) {
					if pattern[pi] == uint8(c) {
						pi++
					}
				}
			} else {
				if pi < len(pattern) && pattern[pi] == uint8(c) {
					pi++
				} else {
					return false
				}
			}
		}
		if pi == len(pattern) {
			return true
		}
		return false
	}
	res := make([]bool, len(queries))
	for i := 0; i < len(res); i++ {
		res[i] = match(queries[i])
	}
	return res
}
