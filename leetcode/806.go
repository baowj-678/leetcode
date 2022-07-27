/**
 * @author: baowj
 * @time: 2022/4/12 21:34:47
 */
package main

func numberOfLines(widths []int, s string) []int {
	lines := 0
	cnt := 0
	for _, c := range s {
		c -= 'a'
		if cnt+widths[c] > 100 {
			lines++
			cnt = widths[c]
		} else {
			cnt += widths[c]
		}
	}
	if cnt > 0 {
		lines++
	}
	return []int{lines, cnt}
}
