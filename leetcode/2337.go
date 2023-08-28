/**
 * @author: baowj
 * @time: 2023/8/28 11:15
 */
package main

func canChange(start string, target string) bool {
	l, r := 0, 0
	for i := 0; i < len(start); i++ {
		switch start[i] {
		case 'L':
			l++
		case 'R':
			r++
		}
		switch target[i] {
		case 'L':
			l--
		case 'R':
			r--
		}
	}
	if l != 0 || r != 0 {
		return false
	}
	for i, j := 0, 0; i < len(target); i++ {
		switch target[i] {
		case 'L':
			for j < len(start) && start[j] != 'L' {
				j++
			}
			if j < len(start) && start[j] == 'L' {
				if j < i {
					return false
				} else {
					j++
				}
			} else {
				return false
			}
		case 'R':
			for j < len(start) && start[j] != 'R' {
				j++
			}
			if j < len(start) && start[j] == 'R' {
				if j > i {
					return false
				} else {
					j++
				}
			} else {
				return false
			}
		}
	}
	return true
}
