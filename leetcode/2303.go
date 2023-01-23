/**
 * @author: baowj
 * @time: 2023/1/23 16:36:42
 * @email: bwj_678@qq.com
 */
package main

func calculateTax(brackets [][]int, income int) float64 {
	res := float64(0)
	addUp := 0
	for _, bracket := range brackets {
		if income > bracket[0] {
			res += float64(bracket[0]-addUp) * float64(bracket[1]) / 100
		} else {
			res += float64(income-addUp) * float64(bracket[1]) / 100
			break
		}
		addUp = bracket[0]
	}
	return res
}
