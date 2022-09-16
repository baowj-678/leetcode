/**
 * @author: baowj
 * @time: 2022/8/13 10:00:08
 */
package main

func maxChunksToSorted(arr []int) int {
	flag := make([]bool, len(arr))
	for i := 0; i < len(flag); i++ {
		flag[i] = false
	}
	cnt := 0
	checkAllTrue := func(flag []bool, idx int) bool {
		for i := 0; i <= idx; i++ {
			if flag[i] == false {
				return false
			}
		}
		return true
	}
	for i, x := range arr {
		flag[x] = true
		if checkAllTrue(flag, i) {
			cnt += 1
		}
	}
	return cnt
}
