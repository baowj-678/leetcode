/**
 * @author: baowj
 * @time: 2022/10/4 10:04:00
 */
package main

func canTransform(start string, end string) bool {
	if len(start) != len(end) {
		return false
	}
	i, j := 0, 0
	for {
		for i < len(start) && start[i] == 'X' {
			i++
		}
		for j < len(end) && end[j] == 'X' {
			j++
		}
		if i < len(start) && j < len(end) {
			if start[i] == end[j] {
				if start[i] == 'L' && i <= j {
					i++
					j++
				} else if start[i] == 'R' && i >= j {
					i++
					j++
				} else {
					return false
				}
			} else {
				return false
			}
		} else {
			break
		}
	}
	if i == len(start) && j == len(end) {
		return true
	}
	return false
}
