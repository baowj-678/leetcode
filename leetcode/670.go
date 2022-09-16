/**
 * @author: baowj
 * @time: 2022/9/13 09:21:35
 */
package main

import "strconv"

func maximumSwap(num int) int {
	s := strconv.FormatInt(int64(num), 10)
	res := s
	for i := 0; i < len(s); i++ {
		var max uint8 = 0
		var maxIdx int = 0
		for j := i + 1; j < len(s); j++ {
			if s[j] > max {
				max, maxIdx = s[j], j
			}
		}
		if max > s[i] {
			tmp := make([]byte, len(s))
			for k := 0; k < len(tmp); k++ {
				tmp[k] = s[k]
			}
			tmp[i], tmp[maxIdx] = tmp[maxIdx], tmp[i]
			tmpS := string(tmp)
			if tmpS > res {
				res = tmpS
			}
		}
	}
	ret, _ := strconv.Atoi(res)
	return ret
}
