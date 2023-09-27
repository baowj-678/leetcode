/**
 * @author: baowj
 * @time: 2023/9/22 11:04
 */
package main

func distMoney(money int, children int) int {
	if money == 4 && children == 1 {
		return -1
	}
	money -= children
	if money < 0 {
		return -1
	}
	num := money / 7
	if num > children {
		return children - 1
	}
	if num == children {
		if money-num*7 == 0 {
			return children
		}
		return children - 1
	}
	if money-num*7 == 3 && children-num == 1 {
		return num - 1
	}
	return num
}
