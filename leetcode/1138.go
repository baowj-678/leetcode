/**
 * @author: baowj
 * @time: 2023/2/12 12:00:19
 * @email: bwj_678@qq.com
 */
package main

func alphabetBoardPath(target string) string {
	res := ""
	x, y := 0, 0
	for _, c := range target {
		tx, ty := int(c-'a')/5, int(c-'a')%5
		diffx, diffy := tx-x, ty-y
		xpath, ypath := "D", "R"
		if diffx < 0 {
			xpath = "U"
			diffx = -diffy
		}
		if diffy < 0 {
			ypath = "L"
			diffy = -diffy
		}
		if c == 'z' {
			for i := 0; i < diffy; i++ {
				res += ypath
			}
			for i := 0; i < diffx; i++ {
				res += xpath
			}
		} else {
			for i := 0; i < diffx; i++ {
				res += xpath
			}
			for i := 0; i < diffy; i++ {
				res += ypath
			}
		}
		res += "!"
		x, y = tx, ty
	}
	return res
}
