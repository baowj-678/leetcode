/**
 * @author: baowj
 * @time: 2022/8/10 10:24:08
 */
package main

import "strconv"

func solveEquation(equation string) string {
	numLeft, numRight := 0, 0
	lastIndex, plus, containX := 0, true, false
	i := 0
	for i = 0; i < len(equation); i++ {
		if equation[i] == '-' || equation[i] == '+' {
			if lastIndex == i {
				if equation[i] == '-' {
					plus = false
				} else {
					plus = true
				}
				lastIndex += 1
			} else {
				if !containX {
					num, _ := strconv.Atoi(equation[lastIndex:i])
					if plus {
						num = -num
					}
					numRight += num
				} else {
					var num = 0
					if i-1 == lastIndex {
						num = 1
					} else {
						num, _ = strconv.Atoi(equation[lastIndex : i-1])
					}
					if !plus {
						num = -num
					}
					numLeft += num
				}
				lastIndex = i
				i -= 1
				containX = false
			}
		} else if equation[i] == '=' {
			if !containX {
				num, _ := strconv.Atoi(equation[lastIndex:i])
				if plus {
					num = -num
				}
				numRight += num
			} else {
				var num = 0
				if i-1 == lastIndex {
					num = 1
				} else {
					num, _ = strconv.Atoi(equation[lastIndex : i-1])
				}
				if !plus {
					num = -num
				}
				numLeft += num
			}
			containX = false
			break
		} else if equation[i] == 'x' {
			containX = true
		}
	}
	plus = true
	containX = false
	lastIndex = i + 1
	for i += 1; i < len(equation); i++ {
		if equation[i] == '-' || equation[i] == '+' {
			if lastIndex == i {
				if equation[i] == '-' {
					plus = false
				} else {
					plus = true
				}
				lastIndex += 1
			} else {
				if !containX {
					num, _ := strconv.Atoi(equation[lastIndex:i])
					if !plus {
						num = -num
					}
					numRight += num
				} else {
					var num = 0
					if i-1 == lastIndex {
						num = 1
					} else {
						num, _ = strconv.Atoi(equation[lastIndex : i-1])
					}
					if plus {
						num = -num
					}
					numLeft += num
				}
				lastIndex = i
				i -= 1
				containX = false
			}
		} else if equation[i] == 'x' {
			containX = true
		}
	}
	if !containX {
		num, _ := strconv.Atoi(equation[lastIndex:i])
		if !plus {
			num = -num
		}
		numRight += num
	} else {
		var num = 0
		if i-1 == lastIndex {
			num = 1
		} else {
			num, _ = strconv.Atoi(equation[lastIndex : i-1])
		}
		if plus {
			num = -num
		}
		numLeft += num
	}

	if numLeft == 0 {
		if numRight == 0 {
			return "Infinite solutions"
		} else {
			return "No solution"
		}
	} else {
		return "x=" + strconv.Itoa(numRight/numLeft)
	}
}

func main() {
	a := "x+5-3+x=6+x-2"
	println(solveEquation(a))
}
