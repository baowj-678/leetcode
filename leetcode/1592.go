/**
 * @author: baowj
 * @time: 2022/9/7 09:24:17
 */
package main

func reorderSpaces(text string) string {
	res := make([]byte, len(text))
	spaceCnt := 0
	wordIdx := []int{}
	isWord := false
	for i, c := range text {
		if c == ' ' {
			spaceCnt++
			if isWord {
				isWord = false
			}
		} else {
			if !isWord {
				wordIdx = append(wordIdx, i)
			}
			isWord = true
		}
	}
	p := 0
	for _, idx := range wordIdx[:len(wordIdx)-1] {
		for text[idx] != ' ' {
			res[p] = text[idx]
			p++
			idx++
		}
		for i := 0; i < spaceCnt/(len(wordIdx)-1); i++ {
			res[p] = ' '
			p++
		}
	}
	idx := wordIdx[len(wordIdx)-1]
	for idx < len(text) && text[idx] != ' ' {
		res[p] = text[idx]
		idx++
		p++
	}
	if len(wordIdx) > 1 {
		for i := 0; i < spaceCnt%(len(wordIdx)-1); i++ {
			res[p] = ' '
			p++
		}
	} else {
		for i := 0; i < spaceCnt; i++ {
			res[p] = ' '
			p++
		}
	}
	return string(res)
}

func main() {
	s := " practice   makes   perfect"
	println(reorderSpaces(s))
}
