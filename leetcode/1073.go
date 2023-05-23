/**
 * @author: baowj
 * @time: 2023/5/18 9:54
 */
package main

func addNegabinary(arr1 []int, arr2 []int) []int {
	rev := func(arr []int) []int {
		for i, j := 0, len(arr)-1; i < j; i, j = i+1, j-1 {
			arr[i], arr[j] = arr[j], arr[i]
		}
		return arr
	}

	arr1 = rev(arr1)
	arr2 = rev(arr2)
	res := make([]int, len(arr1)+len(arr2)+10)
	length := len(arr1)
	if len(arr2) > length {
		length = len(arr2)
	}
	lastOne := 0
	for i := 0; i < len(res); i++ {
		v := res[i]
		if i < len(arr1) {
			v += arr1[i]
		}
		if i < len(arr2) {
			v += arr2[i]
		}
		if v >= 2 {
			if i+1 < len(res) && res[i+1] >= 1 {
				res[i+1]--
				v -= 2
			} else {
				res[i+2]++
				res[i+1]++
				v -= 2
			}
		}
		if v == 1 {
			res[i] = 1
		} else {
			res[i] = 0
		}
		if res[i] == 1 {
			lastOne = i
		}
	}
	res = res[:lastOne+1]
	res = rev(res)
	return res
}

//-32  16 -8  4 -2  1
//         1  0  0  1   -7
//         1  0  1  1   -9
//         0  0  0  0
