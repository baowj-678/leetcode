package main

func minimumSeconds(nums []int) int {
	numIdx := map[int][]int{}
	for i, num := range nums {
		if v, ok := numIdx[num]; ok {
			numIdx[num] = append(v, i)
		} else {
			numIdx[num] = []int{i}
		}
	}
	res := len(nums)
	for _, v := range numIdx {
		max := 0
		if len(v) == 1 {
			max = len(nums) - 1
		} else {
			for i := 0; i < len(v)-1; i++ {
				if v[i+1]-v[i]-1 > max {
					max = v[i+1] - v[i] - 1
				}
			}
			if v[0]+len(nums)-v[len(v)-1]-1 > max {
				max = v[0] + len(nums) - v[len(v)-1] - 1
			}
		}

		max = (max + 1) / 2
		if max < res {
			res = max
		}
	}

	return res
}

// 0 0 1
// 0 0 0 1
// 0 1 0 0 1 0
