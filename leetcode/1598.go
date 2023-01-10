/**
 * @author: baowj
 * @time: 2022/9/9 10:03:30
 */
package main

func minOperations1598(logs []string) int {
	res := 0
	for _, log := range logs {
		if log[0] != '.' {
			res++
		} else if log == "../" {
			if res > 0 {
				res--
			}
		}
	}
	return res
}
