/**
 * @author: baowj
 * @time: 2022/8/14 10:45:51
 */
package main

import (
	"strconv"
)

func smallestNumber(pattern string) string {
	flag := make([]bool, 9)
	for i := 0; i < len(flag); i++ {
		flag[i] = false
	}
	for i := 0; i < len(flag); i++ {
		flag[i] = true
		res := dfs(flag, pattern, 0, i)
		if res != "" {
			return res
		}
		flag[i] = false
	}
	return ""
}

func dfs(flag []bool, pattern string, idx int, last int) string {
	for i := 0; i < len(flag); i++ {
		if flag[i] == true {
			continue
		}
		if pattern[idx] == 'I' && i > last {
			flag[i] = true
			var res string
			if idx+1 == len(pattern) {
				res = strconv.Itoa(i + 1)
			} else {
				res = dfs(flag, pattern, idx+1, i)
			}
			flag[i] = false
			if res != "" {
				return strconv.Itoa(last+1) + res
			}
		} else if pattern[idx] == 'D' && i < last {
			flag[i] = true
			var res string
			if idx+1 == len(pattern) {
				res = strconv.Itoa(i + 1)
			} else {
				res = dfs(flag, pattern, idx+1, i)
			}
			flag[i] = false
			if res != "" {
				return strconv.Itoa(last+1) + res
			}
		}
	}
	return ""
}

func main() {
	print(smallestNumber("DDD"))
}
