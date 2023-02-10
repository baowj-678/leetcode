/**
 * @author: baowj
 * @time: 2022/11/16 23:32:34
 * @email: bwj_678@qq.com
 */
package main

func splitArraySameAverage(nums []int) bool {
	if len(nums) == 1 {
		return false
	}
	sum := 0
	for _, num := range nums {
		sum += num
	}
	for i := 0; i < len(nums); i++ {
		nums[i] = nums[i]*len(nums) - sum
	}
	ma := make(map[int]struct{})
	mb := make(map[int]struct{})
	dfs805(nums[:len(nums)/2], 0, 0, ma, 0)
	dfs805(nums[len(nums)/2:], 0, 0, mb, 0)

	suma := 0
	for _, num := range nums[:len(nums)/2] {
		suma += num
	}
	if suma == 0 {
		return true
	}
	for key, _ := range ma {
		if key == 0 {
			return true
		}
		if _, ok := mb[-key]; ok {
			return true
		}
	}

	for key, _ := range mb {
		if key == 0 {
			return true
		}
		if _, ok := ma[-key]; ok {
			return true
		}
	}
	return false
}

func dfs805(nums []int, idx int, sum int, m map[int]struct{}, record int) {
	if idx == len(nums) {
		if record != (1<<len(nums))-1 {
			if record != 0 {
				m[sum] = struct{}{}
			}
		}
		return
	}
	dfs805(nums, idx+1, sum, m, record)
	dfs805(nums, idx+1, sum+nums[idx], m, record|0b1<<idx)
}
