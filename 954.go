/**
 * @author: baowj
 * @time: 2022/4/1 21:12:08
 */
package main

import "sort"

func canReorderDoubled(arr []int) bool {
	sort.Ints(arr)
	const INT_MAX = int((^uint(0)) >> 1)
	j := 0
	i := 0
	for i = 0; j < len(arr) && i < len(arr); {
		if arr[j] >= 0 {
			break
		}
		if arr[i] == INT_MAX {
			i++
			continue
		}
		if arr[j] == INT_MAX {
			j++
			continue
		}
		if arr[i] == arr[j]*2 {
			arr[i] = INT_MAX
			arr[j] = INT_MAX
			i++
			j++
		} else if arr[i] < arr[j]*2 {
			return false
		} else {
			j++
		}
	}
	for ; i < j; i++ {
		if arr[i] != INT_MAX {
			return false
		}
	}
	i = j
	for j++; j < len(arr) && i < len(arr); {
		if i == j {
			j++
			continue
		}
		if arr[i] == INT_MAX {
			i++
			continue
		}
		if arr[j] == INT_MAX {
			j++
			continue
		}
		if arr[j] == arr[i]*2 {
			arr[i] = INT_MAX
			arr[j] = INT_MAX
			i++
			j++
		} else if arr[i]*2 < arr[j] {
			return false
		} else {
			j++
		}
	}
	for ; i < len(arr); i++ {
		if arr[i] != INT_MAX {
			return false
		}
	}
	return true
}

func main() {
	arr := []int{-2, -4, -8, -8, 26, 13, -16, -4}
	print(canReorderDoubled(arr))
}
