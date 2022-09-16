/**
 * @author: baowj
 * @time: 2022/9/8 22:13:44
 */
package main

import (
	"sort"
)

func maxNumberOfFamilies(n int, reservedSeats [][]int) int {
	sort.Slice(reservedSeats, func(i, j int) bool {
		if reservedSeats[i][0] < reservedSeats[j][0] {
			return true
		} else if reservedSeats[i][0] > reservedSeats[j][0] {
			return false
		} else {
			return reservedSeats[i][1] < reservedSeats[j][1]
		}
	})
	res := 0
	lastRow := 1
	seats := [10]bool{}
	for _, seat := range reservedSeats {
		if seat[0] == lastRow {
			seats[seat[1]-1] = true
		} else {
			res += rowNumbersOfFamilies(seats)
			if seat[0] > lastRow+1 {
				res += (seat[0] - lastRow - 1) * 2
			}
			lastRow = seat[0]
			for i := 0; i < 10; i++ {
				seats[i] = false
			}
			seats[seat[1]-1] = true
		}
	}
	res += rowNumbersOfFamilies(seats)
	res += (n - lastRow) * 2
	return res
}

func rowNumbersOfFamilies(seats [10]bool) int {
	a := !seats[1] && !seats[2]
	b := !seats[3] && !seats[4]
	c := !seats[5] && !seats[6]
	d := !seats[7] && !seats[8]
	if a && b && c && d {
		return 2
	}
	if a && b {
		return 1
	}
	if b && c {
		return 1
	}
	if c && d {
		return 1
	}
	return 0
}
