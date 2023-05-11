/**
 * @author: baowj
 * @time: 2023/5/6 9:53
 */
package main

func minNumberOfFrogs(croakOfFrogs string) int {
	croakNum := [26]int{}
	res := 0
	sum := 0
	for _, c := range croakOfFrogs {
		switch c {
		case 'c':
			croakNum['c'-'a']++
			sum++
		case 'r':
			if croakNum['c'-'a'] > 0 {
				croakNum['c'-'a']--
				croakNum['r'-'a']++
			} else {
				return -1
			}
		case 'o':
			if croakNum['r'-'a'] > 0 {
				croakNum['r'-'a']--
				croakNum['o'-'a']++
			} else {
				return -1
			}
		case 'a':
			if croakNum['o'-'a'] > 0 {
				croakNum['o'-'a']--
				croakNum[0]++
			} else {
				return -1
			}
		case 'k':
			if croakNum[0] > 0 {
				croakNum[0]--
				sum--
			} else {
				return -1
			}
		}

		if res < sum {
			res = sum
		}
	}

	if croakNum['c'-'a'] == 0 && croakNum['r'-'a'] == 0 && croakNum['o'-'a'] == 0 && croakNum[0] == 0 {
		return res
	} else {
		return -1
	}
}
