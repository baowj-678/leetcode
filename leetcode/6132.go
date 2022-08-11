/**
 * @author: baowj
 * @time: 2022/7/31 10:31:49
 */
package main

import "sort"

func minimumOperations(nums []int) int {
	sort.Ints(nums)
	last := -1
	cnt := 0
	for i := 0; i < len(nums); i++ {
		if nums[i] > 0 && last != nums[i] {
			cnt += 1
		}
		last = nums[i]
	}
	return cnt
}

func main() {

}
