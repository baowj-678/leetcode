/**
 * @author: baowj
 * @time: 2022/7/17 10:23:01
 */
package main

func arrayNesting(nums []int) int {
	max := 0
	for i := 0; i < len(nums); i++ {
		if nums[i] != -1 {
			sum := dfs565(nums, i)
			if sum > max {
				max = sum
			}
		}
	}
	return max
}

func dfs565(nums []int, index int) int {
	if nums[index] == -1 {
		return 0
	}
	num := nums[index]
	nums[index] = -1
	num = dfs565(nums, num)
	num += 1
	return num
}

func main() {
	nums := []int{5, 4, 0, 3, 1, 6, 2}
	println(arrayNesting(nums))
}
