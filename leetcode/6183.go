/**
 * @author: baowj
 * @time: 2022/9/18 15:20:41
 */
package main

type node6183 struct {
	alpha int32
	times int
	nexts map[int32]*node6183
}

func sumPrefixScores(words []string) []int {
	// build prefix tree
	root := &node6183{
		nexts: map[int32]*node6183{},
	}
	for _, word := range words {
		root.addPrefix(word)
	}
	//
	res := make([]int, len(words))
	for i, word := range words {
		res[i] = root.getScore(word)
	}
	return res
}

func (this *node6183) addPrefix(word string) {
	for _, c := range word {
		if tmp, ok := this.nexts[c]; ok {
			this = tmp
			this.times++
		} else {
			this.nexts[c] = &node6183{
				alpha: c,
				times: 1,
				nexts: map[int32]*node6183{},
			}
			this = this.nexts[c]
		}
	}
}

func (this *node6183) getScore(word string) int {
	res := 0
	for _, c := range word {
		this = this.nexts[c]
		res += this.times
	}
	return res
}
