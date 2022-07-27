/**
 * @author: baowj
 * @time: 2022/4/7 22:40:45
 */
package main

import "strings"

func rotateString(s string, goal string) bool {
	if len(s) != len(goal) {
		return false
	}
	ss := s + s
	return strings.Contains(ss, goal)
}
