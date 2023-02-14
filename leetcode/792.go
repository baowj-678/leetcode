/**
 * @author: baowj
 * @time: 2022/11/17 11:19
 */
package main

func numMatchingSubseq(s string, words []string) int {
	idxs := [26][]int{}
	for i := 0; i < 26; i++ {
		idxs[i] = []int{}
	}
	for i, c := range s {
		idxs[c-'a'] = append(idxs[c-'a'], i)
	}
	res := 0
	for _, word := range words {
		if f792(word, idxs) {
			res++
		}
	}
	return res
}

func f792(word string, idxs [26][]int) bool {
	i := -1
	for _, c := range word {
		idx := idxs[c-'a']
		left, right := 0, len(idx)-1
		for left < right {
			mid := (left + right) / 2
			if idx[mid] <= i {
				left = mid + 1
			} else {
				right = mid
			}
		}
		if left < len(idx) && idx[left] > i {
			i = idx[left]
			continue
		} else {
			return false
		}
	}
	return true
}

func main() {
	s := "rwpddkvbnnuglnagtva"
	word := []string{"wpddkvbnn", "lnagtva"}
	print(numMatchingSubseq(s, word))
}
