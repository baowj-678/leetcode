/**
 * @author: baowj
 * @time: 2023/1/23 16:52:41
 * @email: bwj_678@qq.com
 */
package main

func strongPasswordCheckerII(password string) bool {
	if len(password) < 8 {
		return false
	}
	smallAlpha, bigAlpha, hasNum, hasSpecialAlpha := false, false, false, false
	lastC := int32(-1)
	specialAlpha := "!@#$%^&*()-+"
	for _, c := range password {
		if c >= 'a' && c <= 'z' {
			smallAlpha = true
		}
		if c >= 'A' && c <= 'Z' {
			bigAlpha = true
		}
		if c >= '0' && c <= '9' {
			hasNum = true
		}
		for _, sc := range specialAlpha {
			if sc == c {
				hasSpecialAlpha = true
			}
		}
		if lastC == c {
			return false
		}
		lastC = c
	}
	if smallAlpha && bigAlpha && hasNum && hasSpecialAlpha {
		return true
	}
	return false
}
