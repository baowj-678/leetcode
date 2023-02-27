/**
 * @author: baowj
 * @time: 2023/2/27 14:01
 */
package main

func movesToMakeZigzag(nums []int) int {
	even := 0
	copyNums := make([]int, len(nums))
	copy(copyNums, nums)
	for i := 0; i < len(copyNums); i += 2 {
		if i-1 >= 0 && copyNums[i] <= copyNums[i-1] {
			even += copyNums[i-1] + 1 - copyNums[i]
			copyNums[i-1] = copyNums[i] - 1
		}
		if i+1 < len(copyNums) && copyNums[i] <= copyNums[i+1] {
			even += copyNums[i+1] + 1 - copyNums[i]
			copyNums[i+1] = copyNums[i] - 1
		}
	}
	odd := 0
	for i := 1; i < len(nums); i += 2 {
		if i-1 >= 0 && nums[i] <= nums[i-1] {
			odd += nums[i-1] + 1 - nums[i]
			nums[i-1] = nums[i] - 1
		}
		if i+1 < len(nums) && nums[i] <= nums[i+1] {
			odd += nums[i+1] + 1 - nums[i]
			nums[i+1] = nums[i] - 1
		}
	}

	if even > odd {
		return odd
	}
	return even
}
