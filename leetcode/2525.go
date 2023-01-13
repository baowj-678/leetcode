/**
 * @author: baowj
 * @time: 2023/1/13 11:45:34
 * @email: bwj_678@qq.com
 */
package main

func categorizeBox(length int, width int, height int, mass int) string {
	isBulky := false
	if length >= 1e4 || width >= 1e4 || height >= 1e4 || int64(length)*int64(width)*int64(height) >= 1e9 {
		isBulky = true
	}
	isHeavy := false
	if mass >= 100 {
		isHeavy = true
	}
	if isHeavy && isBulky {
		return "Both"
	} else if isHeavy {
		return "Heavy"
	} else if isBulky {
		return "Bulky"
	} else {
		return "Neither"
	}
}
