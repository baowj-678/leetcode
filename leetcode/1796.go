/**
 * @author: baowj
 * @time: 2022/12/3 12:27:24
 * @email: bwj_678@qq.com
 */
package main

func secondHighest(s string) int {
	max, secondMax := -1, -1
	for _, c := range s {
		c -= '0'
		if c >= 0 && c <= 9 {
			if int(c) > max {
				secondMax = max
				max = int(c)
			} else if int(c) > secondMax && int(c) < max {
				secondMax = int(c)
			}
		}
	}
	return secondMax
}

func main() {
	println(secondHighest("dfa12321afd"))
}
