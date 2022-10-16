/**
 * @author: baowj
 * @time: 2022/10/16 17:57:16
 */
package main

func buildArray(target []int, n int) []string {
	res := []string{}
	for t, i := 0, 1; i <= n && t < len(target); {
		if target[t] == i {
			res = append(res, "Push")
			i++
			t++
		} else if target[t] > i {
			res = append(res, "Push")
			res = append(res, "Pop")
			i++
		}
	}
	return res
}
