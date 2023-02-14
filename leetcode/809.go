/**
 * @author: baowj
 * @time: 2022/11/25 13:49
 */
package main

func expressiveWords(s string, words []string) int {
	res := 0
	alpha := []int32{}
	count := []int32{}
	var lastC int32 = -1
	for _, a := range s {
		if lastC == a {
			count[len(count)-1]++
		} else {
			alpha = append(alpha, a)
			count = append(count, 1)
			lastC = a
		}
	}

	for _, word := range words {
		alpha_ := []int32{}
		count_ := []int32{}
		lastC = -1
		for _, a := range word {
			if lastC == a {
				count_[len(count_)-1]++
			} else {
				alpha_ = append(alpha_, a)
				count_ = append(count_, 1)
				lastC = a
			}
		}
		if len(alpha) != len(alpha_) {
			continue
		}

		isOK := true
		for i := 0; i < len(alpha); i++ {
			if alpha[i] != alpha_[i] {
				isOK = false
				break
			}
			if !(count_[i] == count[i] || (count_[i] < count[i] && count[i] >= 3)) {
				isOK = false
				break
			}
		}
		if isOK {
			res++
		}
	}
	return res
}

func main() {
	s := "heeellooo"
	words := []string{"hello", "hi", "helo"}
	print(expressiveWords(s, words))
}
