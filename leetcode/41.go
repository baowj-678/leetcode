/**
 * @author: baowj
 * @time: 2022/8/26 21:30:31
 */
package main

func firstMissingPositive(nums []int) int {
	for i := 0; i < len(nums); i++ {
		j := nums[i] - 1
		if j != i {
			nums[i] = -1
			for j >= 0 && j < len(nums) {
				if j == nums[j]-1 {
					break
				}
				nums[j], j = j+1, nums[j]-1
			}
		}
	}
	for i := 0; i < len(nums); i++ {
		if nums[i] != i+1 {
			return i + 1
		}
	}
	return len(nums) + 1
}

func main() {
	arr := []int{3, 4, -1, 1}
	firstMissingPositive(arr)
}
