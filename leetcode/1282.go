/**
 * @author: baowj
 * @time: 2022/8/12 10:44:30
 */
package main

func groupThePeople(groupSizes []int) [][]int {
	sizeGroupMap := map[int][]int{}
	ret := [][]int{}
	for i, gSize := range groupSizes {
		if gSize == 1 {
			ret = append(ret, []int{i})
		} else {
			if value, ok := sizeGroupMap[gSize]; ok {
				value = append(value, i)
				if len(value) == gSize {
					ret = append(ret, value)
					value = []int{}
				}
				sizeGroupMap[gSize] = value
			} else {
				value = []int{i}
				sizeGroupMap[gSize] = value
			}
		}
	}
	return ret
}
