/**
 * @author: baowj
 * @time: 2022/12/15 22:17:42
 * @email: bwj_678@qq.com
 */
package main

func checkIfPangram(sentence string) bool {
	m := map[int]struct{}{}
	for _, c := range sentence {
		m[int(c)] = struct{}{}
	}
	return len(m) == 26
}
