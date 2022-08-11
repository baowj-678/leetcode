/**
 * @author: baowj
 * @time: 2022/8/1 09:55:19
 */
package main

func generateTheString(n int) string {
	res := ""
	if n%2 == 0 {
		res += "a"
		for i := 0; i < n-1; i++ {
			res += "b"
		}
	} else {
		for i := 0; i < n; i++ {
			res += "a"
		}
	}
	return res
}
