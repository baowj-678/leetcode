/**
 * @author: baowj
 * @time: 2022/4/17 10:56:22
 */
package main

func largestPalindrome(n int) int {
	if n == 1 {
		return 9
	}
	lower := geten(n-1) - 1
	upper := geten(n) - 1
	for left := upper; left > lower; left-- {
		p := int64(left)
		for x := int64(left); x > 0; x /= 10 {
			p = p*10 + x%10
		}
		for x := int64(upper); x*x >= p; x-- {
			if p%x == 0 {
				return int(p % 1337)
			}
		}
	}
	return 0
}

func geten(n int) int {
	v := 1
	for i := 0; i < n; i++ {
		v *= 10
	}
	return v
}

func main() {
	println(largestPalindrome(8))
}
