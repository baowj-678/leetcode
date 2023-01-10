/**
 * @author: baowj
 * @time: 2023/1/8 11:42:25
 * @email: bwj_678@qq.com
 */
package main

func repeatedCharacter(s string) byte {
	cnt := make([]int, 26)
	for _, c := range s {
		c -= 'a'
		cnt[c]++
		if cnt[int(c)] == 2 {
			return byte(c + 'a')
		}
	}
	return 0
}
