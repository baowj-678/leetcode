/**
 * @author: baowj
 * @time: 2022/4/22 09:50:17
 */
package main

func maxRotateFunction(nums []int) int {
	maxValue := 0
	sum := 0
	for i := 0; i < len(nums); i++ {
		sum += nums[i]
		maxValue += i * nums[i]
	}
	tmpValue := maxValue
	for i := len(nums) - 1; i > 0; i-- {
		tmpValue = tmpValue + sum - len(nums)*nums[i]
		if tmpValue > maxValue {
			maxValue = tmpValue
		}
	}
	return maxValue
}
