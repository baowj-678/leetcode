/**
 * @author: baowj
 * @time: 2022/4/3 09:57:29
 */
package main

func nextGreatestLetter(letters []byte, target byte) byte {
	left := 0
	right := len(letters) - 1
	if letters[right] <= target {
		return letters[left]
	}
	for left < right {
		mid := (left + right) / 2
		if letters[mid] > target {
			right = mid
		} else {
			left = mid + 1
		}
	}
	return letters[right]
}

func main() {
	letters := []byte{'c', 'f', 'j'}
	target := byte('z')
	print(string(nextGreatestLetter(letters, target)))
}
