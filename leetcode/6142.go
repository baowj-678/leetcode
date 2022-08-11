/**
 * @author: baowj
 * @time: 2022/8/7 10:56:24
 */
package main

func countBadPairs(nums []int) int64 {
	m := map[int]int{}
	for i := 0; i < len(nums); i++ {
		if _, ok := m[nums[i]-i]; ok {
			m[nums[i]-i] += 1
		} else {
			m[nums[i]-i] = 1
		}
	}
	res := C(len(nums))
	for _, v := range m {
		res -= C(v)
	}
	return res
}

func C(x int) int64 {
	return int64(x) * (int64(x) + 1) / 2
}
