/**
 * @author: baowj
 * @time: 2023/1/11 16:39:55
 * @email: bwj_678@qq.com
 */
package main

func digitCount(num string) bool {
	cnt := [10]int{}
	for _, c := range num {
		cnt[c-'0']++
	}

	for i, c := range num {
		if cnt[i] != int(c-'0') {
			return false
		}
	}
	return true
}
