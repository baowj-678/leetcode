/**
 * @author: baowj
 * @time: 2023/3/24 15:43
 */
package main

type StreamChecker struct {
	m       map[string]struct{}
	length  int
	letters []byte
	idx     int
}

func Constructor1032(words []string) StreamChecker {
	max := func(a, b int) int {
		if a > b {
			return a
		}
		return b
	}
	m := map[string]struct{}{}
	length := 0
	for _, word := range words {
		length = max(length, len(word))
		m[word] = struct{}{}
	}
	return StreamChecker{
		m:       m,
		length:  length,
		letters: make([]byte, 4*1e4+1000),
		idx:     0,
	}
}

func (this *StreamChecker) Query(letter byte) bool {
	this.letters[this.idx] = letter
	this.idx++
	for i := 1; i <= this.length && i <= this.idx; i++ {
		if _, ok := this.m[string(this.letters[this.idx-i:this.idx])]; ok {
			return true
		}
	}
	return false
}

/**
 * Your StreamChecker object will be instantiated and called as such:
 * obj := Constructor(words);
 * param_1 := obj.Query(letter);
 */
