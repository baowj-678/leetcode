package main

import "fmt"

func splitMessage(message string, limit int) []string {
	i, cap := 0, 0
	for {
		i++
		var sLen int
		if i < 10 {
			sLen = 5
		} else if i < 100 {
			if i == 10 {
				cap -= 9
			}
			sLen = 7
		} else if i < 1000 {
			if i == 100 {
				cap -= 99
			}
			sLen = 9
		} else if i < 10000 {
			if i == 1000 {
				cap -= 999
			}
			sLen = 11
		}

		if limit-sLen <= 0 {
			return []string{}
		}
		cap += limit - sLen
		if cap < len(message) {
			continue
		}

		res := []string{}
		j, idx := 0, 1
		for j < len(message) {
			appendS := fmt.Sprintf("<%v/%v>", idx, i)
			var s string
			if j+limit-len(appendS) < len(message) {
				s = message[j:j+limit-len(appendS)] + appendS
			} else {
				s = message[j:] + appendS
			}
			res = append(res, s)
			j = j + limit - len(appendS)
			idx++
		}
		return res
	}
}
