/**
 * @author: baowj
 * @time: 2022/7/27 09:53:41
 */
package main

import (
	"fmt"
	"strconv"
)

func fractionAddition(expression string) string {
	x, y, prevIndex := 0, 1, 0
	for i := 1; i < len(expression); i++ {
		if expression[i] == '+' || expression[i] == '-' {
			numerator, denominator := string2int(expression[prevIndex:i])
			x, y = add(x, y, numerator, denominator)
			prevIndex = i
		}
	}
	numerator, denominator := string2int(expression[prevIndex:])
	x, y = add(x, y, numerator, denominator)
	return strconv.Itoa(x) + "/" + strconv.Itoa(y)
}

func string2int(expression string) (x, y int) {
	for i := 0; i < len(expression); i++ {
		if expression[i] == '/' {
			numerator, err := strconv.Atoi(expression[:i])
			if err != nil {
				err.Error()
			}
			denominator, err := strconv.Atoi(expression[i+1:])
			if err != nil {
				err.Error()
			}
			return numerator, denominator
		}
	}
	return 0, 1
}

func add(xNumerator, xDenominator, yNumerator, yDenominator int) (numerator, denominator int) {
	gcdVal := gcd(xDenominator, yDenominator)
	x := xDenominator * yDenominator / gcdVal / xDenominator
	y := xDenominator * yDenominator / gcdVal / yDenominator
	yNumerator *= y
	yDenominator *= y
	xNumerator *= x
	xDenominator *= x
	xNumerator += yNumerator

	if xNumerator < 0 {
		gcdVal = gcd(-xNumerator, xDenominator)
	} else {
		gcdVal = gcd(xNumerator, xDenominator)
	}
	if gcdVal != 0 {
		xNumerator /= gcdVal
		xDenominator /= gcdVal
	} else {
		if xNumerator == 0 {
			xDenominator = 1
		}
	}
	return xNumerator, xDenominator
}

func gcd(x, y int) int {
	if x == 0 || y == 0 {
		return 0
	}
	if x < y {
		x, y = y, x
	}
	for y != 0 {
		x, y = y, x%y
	}
	return x
}

func main() {
	var s string
	fmt.Scan(&s)
	println(fractionAddition(s))
}
