/**
 * @author: baowj
 * @time: 2022/8/18 10:09:01
 */
package main

func maxEqualFreq(nums []int) int {
	count := map[int]int{}
	freq := map[int]int{}
	maxFreq := 0
	maxLength := 0
	for i, num := range nums {
		if x, ok := count[num]; ok {
			if maxFreq < x+1 {
				maxFreq = x + 1
			}
			count[num] = x + 1
			freq[x]--
			if freq[x] == 0 {
				delete(freq, x)
			}
			if f, ok := freq[x+1]; ok {
				freq[x+1] = f + 1
			} else {
				freq[x+1] = 1
			}
		} else {
			if maxFreq < 1 {
				maxFreq = 1
			}
			count[num] = 1
			if x, ok := freq[1]; ok {
				freq[1] = x + 1

			} else {
				freq[1] = 1
			}
		}
		if len(freq) == 1 {
			if maxFreq == 1 {
				maxLength = i + 1
			} else {
				for _, value := range freq {
					if value == 1 {
						maxLength = i + 1
					}
				}
			}
		} else if len(freq) == 2 {
			if _, ok := freq[maxFreq-1]; ok && freq[maxFreq] == 1 {
				maxLength = i + 1
			} else if v, ok := freq[1]; ok && v == 1 {
				maxLength = i + 1
			}
		}
	}
	return maxLength
}

func main() {
	nums := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	print(maxEqualFreq(nums))
}
