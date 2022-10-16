/**
 * @author: baowj
 * @time: 2022/9/18 10:03:21
 */
package main

import "strconv"

func countDaysTogether(arriveAlice string, leaveAlice string, arriveBob string, leaveBob string) int {
	aaMonth, _ := strconv.Atoi(arriveAlice[:2])
	aaDay, _ := strconv.Atoi(arriveAlice[3:])
	abMonth, _ := strconv.Atoi(arriveBob[:2])
	abDay, _ := strconv.Atoi(arriveBob[3:])
	arrive := [2]int{}
	if aaMonth < abMonth {
		arrive[0], arrive[1] = abMonth, abDay
	} else if aaMonth > abMonth {
		arrive[0], arrive[1] = aaMonth, aaDay
	} else {
		if aaDay < abDay {
			arrive[0], arrive[1] = abMonth, abDay
		} else {
			arrive[0], arrive[1] = aaMonth, aaDay
		}
	}

	laMonth, _ := strconv.Atoi(leaveAlice[:2])
	laDay, _ := strconv.Atoi(leaveAlice[3:])
	lbMonth, _ := strconv.Atoi(leaveBob[:2])
	lbDay, _ := strconv.Atoi(leaveBob[3:])
	leave := [2]int{}
	if laMonth < lbMonth {
		leave[0], leave[1] = laMonth, laDay
	} else if laMonth > lbMonth {
		leave[0], leave[1] = lbMonth, lbDay
	} else {
		if laDay < lbDay {
			leave[0], leave[1] = laMonth, laDay
		} else {
			leave[0], leave[1] = lbMonth, lbDay
		}
	}

	days := []int{0, 31, 28, 31, 30, 31, 30, 31, 31, 30, 31, 30, 31}
	res := 0
	if leave[0] < arrive[0] {
		return 0
	} else if leave[0] == arrive[0] {
		res = leave[1] - arrive[1] + 1
		if res > 0 {
			return res
		} else {
			return 0
		}
	}
	res += days[arrive[0]] - arrive[1] + 1
	res += leave[1]
	for i := arrive[0] + 1; i < leave[0]; i++ {
		res += days[i]
	}
	return res
}
