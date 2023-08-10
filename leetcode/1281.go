package main

func subtractProductAndSum(n int) int {
	sum, product := 0, 1
	for n > 0 {
		tmp := n % 10
		n /= 10
		sum += tmp
		product *= tmp
	}
	return product - sum
}
