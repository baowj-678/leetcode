/**
 * @author: baowj
 * @time: 2022/9/22 09:56:36
 */
package main

func kSimilarity(s1 string, s2 string) int {
	type pair struct {
		s string
		i int
	}
	q := []pair{{s: s1, i: 0}}
	vis := map[string]bool{s1: true}
	step := 0
	for ; ; step++ {
		tmp := q
		q = nil
		for _, p := range tmp {
			s, i := p.s, p.i
			for ; i < len(s2) && s[i] == s2[i]; i++ {
			}
			if i == len(s2) {
				return step
			}
			sbyte := []byte(s)
			for j := i + 1; j < len(s2); j++ {
				if s2[i] == s[j] && s[j] != s2[j] {
					sbyte[i], sbyte[j] = sbyte[j], sbyte[i]
					if snew := string(sbyte); !vis[snew] {
						vis[snew] = true
						q = append(q, pair{snew, i + 1})
					}
					sbyte[i], sbyte[j] = sbyte[j], sbyte[i]
				}
			}
		}
	}
	return step
}
