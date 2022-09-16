/**
 * @author: baowj
 * @time: 2022/9/11 15:51:34
 */
package main

func partitionString(s string) int {
	res := 0
	flag := [26]bool{}
	for _, x := range s {
		x -= 'a'
		if flag[x] {
			res++
			for i := 0; i < 26; i++ {
				flag[i] = false
			}
			flag[x] = true
		} else {
			flag[x] = true
		}
	}
	for _, x := range flag {
		if x {
			res++
			break
		}
	}
	return res
}
