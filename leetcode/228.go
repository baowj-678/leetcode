package main

import (
	"fmt"
	"strconv"
)

func summaryRanges(nums []int) []string {
	var res []string
	if len(nums) == 0 {
		return res
	}
	start, last := nums[0], nums[0]
	for _, num := range nums {
		if num > last+1 {
			if start == last {
				res = append(res, strconv.Itoa(start))
			} else {
				res = append(res, fmt.Sprintf("%s->%s", strconv.Itoa(start), strconv.Itoa(last)))
			}
			start = num
		}
		last = num
	}
	if start == last {
		res = append(res, strconv.Itoa(start))
	} else {
		res = append(res, fmt.Sprintf("%s->%s", strconv.Itoa(start), strconv.Itoa(last)))
	}
	return res
}
