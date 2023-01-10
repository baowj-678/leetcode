/**
 * @author: baowj
 * @time: 2022/11/27 16:23:54
 * @email: bwj_678@qq.com
 */
package main

var MOD6251 int64 = 1e9 + 7

func countPalindromes(s string) int {
	xy := make([][100]int, len(s))
	yx := make([][100]int, len(s))
	cnt := make([]int, 10)
	for i, c := range s {
		c -= '0'
		if i > 0 {
			for k := 0; k < 100; k++ {
				xy[i][k] = xy[i-1][k]
			}
		}
		for x, num := range cnt {
			v := x*10 + int(c)
			xy[i][v] += num
		}
		cnt[int(c)]++
	}

	for i := 0; i < len(cnt); i++ {
		cnt[i] = 0
	}
	for i := len(s) - 1; i >= 0; i-- {
		c := s[i] - '0'
		if i < len(s)-1 {
			for k := 0; k < 100; k++ {
				yx[i][k] = yx[i+1][k]
			}
		}
		for x, num := range cnt {
			v := x*10 + int(c)
			yx[i][v] += num
		}
		cnt[int(c)]++
	}
	var res int64 = 0
	for i := 2; i < len(s)-2; i++ {
		for j := 0; j < 100; j++ {
			res = (res + int64(xy[i-1][j])*int64(yx[i+1][j])) % MOD6251
		}
	}
	return int(res)
}

func main() {
	println(countPalindromes("103301"))
}
