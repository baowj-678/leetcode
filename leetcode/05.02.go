/**
 * @author: baowj
 * @time: 2023/3/4 16:02:45
 * @email: bwj_678@qq.com
 */
package main

func printBin(num float64) string {
	res := "0."
	for num != 0 && len(res) < 32 {
		num *= 2
		if num >= 1.0 {
			res = res + "1"
			num -= 1.0
		} else {
			res = res + "0"
		}
	}
	if num != 0.0 {
		return "ERROR"
	} else {
		return res
	}
}
