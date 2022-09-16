/**
 * @author: baowj
 * @time: 2022/9/11 15:36:11
 */
package main

func mostFrequentEven(nums []int) int {
	m := map[int]int{}
	for _, x := range nums {
		if x%2 == 1 {
			continue
		}
		if v, ok := m[x]; ok {
			m[x] = v + 1
		} else {
			m[x] = 1
		}
	}
	maxKey, maxValue := -1, -1
	for key, value := range m {
		if value > maxValue || (value == maxValue && key < maxKey) {
			maxKey = key
			maxValue = value
		}
	}
	return maxKey
}

func main() {
	nums := []int{0, 1, 2, 2, 4, 4, 1}
	println(mostFrequentEven(nums))
}
