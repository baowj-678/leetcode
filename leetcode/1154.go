package main

import "strconv"

func dayOfYear(date string) int {
	year, _ := strconv.Atoi(date[:4])
	month, _ := strconv.Atoi(date[5:7])
	day, _ := strconv.Atoi(date[8:])
	monthDays := []int{31, 28, 31, 30, 31, 30, 31, 31, 30, 31, 30, 31}
	if year%4 == 0 {
		monthDays[1] = 29
	}
	if year%100 == 0 {
		monthDays[1] = 28
	}
	if year%400 == 0 {
		monthDays[1] = 29
	}
	res := 0
	for i := 0; i < month-1; i++ {
		res += monthDays[i]
	}
	res += day
	return res
}
