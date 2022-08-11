/**
 * @author: baowj
 * @time: 2022/8/7 11:36:12
 */
package main

func minimumReplacement(nums []int) int64 {
	max := nums[len(nums)-1]
	var cnt int64 = 0
	for i := len(nums) - 2; i >= 0; i-- {
		num := nums[i]
		if max < num {
			k := (num + max - 1) / max
			max = num / k
			cnt += int64(k - 1)
		} else {
			max = num
		}
	}
	return cnt
}
func main() {
	nums := []int{1, 2, 3, 4, 5, 74, 2, 2, 3, 4, 5, 2, 1, 4, 6, 7, 7, 8, 90}
	println(minimumReplacement(nums))
}
