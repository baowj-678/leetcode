/**
 * @author: baowj
 * @time: 2022/10/3 09:52:32
 */
package main

func checkOnesSegment(s string) bool {
	for i := 1; i < len(s); i++ {
		if s[i] == '1' && s[i-1] == '0' {
			return false
		}
	}
	return true
}
