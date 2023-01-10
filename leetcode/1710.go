/**
 * @author: baowj
 * @time: 2022/11/16 22:06:43
 * @email: bwj_678@qq.com
 */
package main

import "sort"

func maximumUnits(boxTypes [][]int, truckSize int) int {
	sort.Slice(boxTypes, func(i, j int) bool {
		return boxTypes[i][1] > boxTypes[j][1]
	})
	res := 0
	for _, box := range boxTypes {
		if box[0] <= truckSize {
			truckSize -= box[0]
			res += box[0] * box[1]
		} else {
			res += truckSize * box[1]
			return res
		}
	}
	return res
}

func main() {
	boxTypes := [][]int{{1, 3}, {2, 2}, {3, 1}}
	println(maximumUnits(boxTypes, 4))
}
