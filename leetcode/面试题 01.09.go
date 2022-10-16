/**
 * @author: baowj
 * @time: 2022/9/29 10:33:15
 */
package main

import "strings"

func isFlipedString(s1 string, s2 string) bool {
	return len(s1) == len(s2) && strings.Contains(s2+s2, s1)
}
