/**
 * @author: baowj
 * @time: 2023/4/27 9:50
 */
package main

type Node1048 struct {
	idx   int
	nexts []*Node1048
}

func longestStrChain(words []string) int {
	isBefore := func(word1, word2 string) bool {
		if len(word1) != len(word2)-1 {
			return false
		}
		for i, j := 0, 0; i < len(word1) && j < len(word2); i, j = i+1, j+1 {
			if word1[i] != word2[j] {
				if i < j {
					return false
				} else {
					i--
				}
			}
		}
		return true
	}

	nodes := make([]Node1048, len(words))
	for i := 0; i < len(words); i++ {
		nodes[i].idx = i
		for j := i + 1; j < len(words); j++ {
			if isBefore(words[i], words[j]) {
				nodes[i].nexts = append(nodes[i].nexts, &nodes[j])
			} else if isBefore(words[j], words[i]) {
				nodes[j].nexts = append(nodes[j].nexts, &nodes[i])
			}
		}
	}
	max := func(a, b int) int {
		if a < b {
			a = b
		}
		return a
	}
	res := 0
	for i := 0; i < len(nodes); i++ {
		res = max(res, dfs1048(&nodes[i]))
	}
	return res
}

func dfs1048(head *Node1048) int {
	res := 1
	for _, next := range head.nexts {
		t := dfs1048(next)
		if t+1 > res {
			res = t + 1
		}
	}
	return res
}
