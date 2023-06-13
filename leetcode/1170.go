/**
 * @author: baowj
 * @time: 2023/6/12 16:20
 */
package main

func numSmallerByFrequency(queries []string, words []string) []int {
	f := func(s string) int {
		cnt := [26]int{}
		for _, c := range s {
			cnt[c-'a']++
		}
		for _, x := range cnt {
			if x != 0 {
				return x
			}
		}
		return 0
	}

	table := [2001]int{}
	for i := 0; i < len(words); i++ {
		table[f(words[i])]++
	}
	for i := len(table) - 2; i >= 0; i-- {
		table[i] += table[i+1]
	}
	res := make([]int, len(queries))
	for i := 0; i < len(queries); i++ {
		v := f(queries[i])
		res[i] = table[v+1]
	}
	return res
}
