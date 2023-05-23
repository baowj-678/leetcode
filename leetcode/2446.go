/**
 * @author: baowj
 * @time: 2023/5/17 9:58
 */
package main

import "strconv"

func haveConflict(event1 []string, event2 []string) bool {
	time1 := [2][2]int{}
	time2 := [2][2]int{}
	for i := 0; i < 2; i++ {
		h, err := strconv.Atoi(event1[i][:2])
		if err != nil {
			return false
		}
		m, err := strconv.Atoi(event1[i][3:])
		if err != nil {
			return false
		}
		time1[i] = [2]int{h, m}

		h, err = strconv.Atoi(event2[i][:2])
		if err != nil {
			return false
		}
		m, err = strconv.Atoi(event2[i][3:])
		if err != nil {
			return false
		}
		time2[i] = [2]int{h, m}
	}

	lq := func(t1, t2 [2]int) bool {
		if t1[0] == t2[0] {
			return t1[1] <= t2[1]
		}
		return t1[0] < t2[0]
	}

	if (lq(time2[0], time1[1]) && lq(time1[0], time2[1])) || (lq(time1[0], time2[1]) && lq(time2[0], time1[1])) {
		return true
	}
	return false
}
