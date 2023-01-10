/**
 * @author: baowj
 * @time: 2022/12/18 15:53:32
 * @email: bwj_678@qq.com
 */
package main

import "fmt"

func smallestValue(n int) int {
	values := []int{}
	values = append(values, n)
	lastN := n
	for {
		n = 0
		for len(values) > 0 {
			v := values[0]
			values = values[1:]
			isPrime := true
			for i := 2; i*i <= v; i++ {
				if v%i == 0 {
					values = append(values, i)
					values = append(values, v/i)
					isPrime = false
					break
				}
			}
			if isPrime {
				n += v
			}
		}
		if lastN == n {
			return n
		}
		lastN = n
		values = []int{n}
	}
}

func main() {
	fmt.Println(smallestValue(4))
}
