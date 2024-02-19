package main

func sumIndicesWithKSetBits(nums []int, k int) int {
	res := 0
	for i, num := range nums {
		cnt, tmp := 0, i
		for tmp > 0 {
			if tmp&0b1 > 0 {
				cnt++
			}
			tmp >>= 1
		}
		if cnt == k {
			res += num
		}
	}
	return res
}
