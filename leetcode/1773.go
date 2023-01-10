/**
 * @author: baowj
 * @time: 2022/11/1 16:44:08
 */
package main

func countMatches(items [][]string, ruleKey string, ruleValue string) int {
	res := 0
	idx := 0
	if ruleKey == "type" {
		idx = 0
	} else if ruleKey == "color" {
		idx = 1
	} else {
		idx = 2
	}
	for _, item := range items {
		if item[idx] == ruleValue {
			res++
		}
	}
	return res
}
