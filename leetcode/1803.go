/**
 * @author: baowj
 * @time: 2023/1/5 15:11:36
 * @email: bwj_678@qq.com
 */
package main

const SHIFT int = 15

type trie struct {
	key   int
	value int
	level int
	tries [2]*trie
}

func (t *trie) add(x int) {
	if t.key == (x>>t.level)&1 {
		t.value++
	} else {
		panic("Error")
	}
	if t.level == 0 {
		return
	}
	nextLevel := t.level - 1
	nextBit := (x >> nextLevel) & 1
	if t.tries[nextBit] == nil {
		t.tries[nextBit] = &trie{
			key:   nextBit,
			value: 0,
			level: nextLevel,
			tries: [2]*trie{nil, nil},
		}
	}
	t.tries[nextBit].add(x)
}

func (t *trie) sum(x int, high int) int {
	if t.level == 0 {
		return t.value
	}
	nextLevel := t.level - 1
	higtBit, xBit := (high>>nextLevel)&1, (x>>nextLevel)&1
	if higtBit == 0 {
		if t.tries[xBit] == nil {
			return 0
		} else {
			return t.tries[xBit].sum(x, high)
		}
	} else {
		s := 0
		if t.tries[xBit] != nil {
			s += t.tries[xBit].value
		}
		if t.tries[1-xBit] != nil {
			s += t.tries[1-xBit].sum(x, high)
		}
		return s
	}
}

func countPairs(nums []int, low int, high int) int {
	t := trie{
		key:   0,
		value: 0,
		level: SHIFT,
		tries: [2]*trie{nil, nil},
	}
	for _, num := range nums {
		t.add(num)
	}
	lowCnt := 0
	for _, num := range nums {
		lowCnt += t.sum(num, low-1)
	}
	highCnt := 0
	for _, num := range nums {
		highCnt += t.sum(num, high)
	}
	return (highCnt - lowCnt) / 2
}

func main() {
	nums := []int{1, 4, 2, 7}
	println(countPairs(nums, 2, 6))
}
