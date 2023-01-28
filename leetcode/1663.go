/**
 * @author: baowj
 * @time: 2023/1/28 11:08:06
 * @email: bwj_678@qq.com
 */
package main

func getSmallestString(n int, k int) string {
	res := make([]byte, n)
	for i := 0; i < n; i++ {
		res[i] = 'a'
	}
	k -= n
	i := n - 1
	for k > 0 && i >= 0 {
		if k >= 25 {
			res[i] = 'z'
			k -= 25
		} else {
			res[i] = 'a' + byte(k)
			k -= k
		}
		i--
	}
	return string(res)
}
