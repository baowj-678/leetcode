/**
 * @author: baowj
 * @time: 2022/9/16 21:11:09
 */
package main

import "sort"

const MOD850 = 1e9 + 7

func rectangleArea(rectangles [][]int) int {
	lines := [][]int{}
	yMap := map[int]int{}
	for _, rectangle := range rectangles {
		lines = append(lines, []int{rectangle[0], rectangle[1], rectangle[3], 0})
		lines = append(lines, []int{rectangle[2], rectangle[1], rectangle[3], 1})
		yMap[rectangle[1]] = 1
		yMap[rectangle[3]] = 1
	}
	yLines := make([]int, len(yMap))
	k := 0
	for key, _ := range yMap {
		yLines[k] = key
		k++
	}
	sort.Slice(lines, func(i, j int) bool {
		return lines[i][0] < lines[j][0]
	})
	sort.Ints(yLines)
	gaps := make([]int, len(yLines))
	res := int64(0)
	prevX := 0
	for _, line := range lines {
		length := getLength(yLines, gaps)
		res = (res + int64(length)*int64(line[0]-prevX)) % MOD850
		if line[3] == 0 {
			changeSegment(yLines, gaps, line[1], line[2], true)
		} else {
			changeSegment(yLines, gaps, line[1], line[2], false)
		}
		prevX = line[0]
	}
	return int(res)
}

func changeSegment(yLines []int, yCounts []int, from int, to int, isAdd bool) {
	for i := 0; i < len(yLines); i++ {
		if yLines[i] >= from && yLines[i] < to {
			if isAdd {
				yCounts[i]++
			} else {
				yCounts[i]--
			}
		}
	}
}

func getLength(yLines []int, yCounts []int) int {
	res := 0
	for i := 0; i < len(yLines); i++ {
		if yCounts[i] > 0 {
			res += yLines[i+1] - yLines[i]
		}
	}
	return res
}
