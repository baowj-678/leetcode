/**
 * @author: baowj
 * @time: 2023/6/14 10:02
 */
package main

func numTimesAllBlue(flips []int) int {
	biggest := 0
	max := func(x, y int) int {
		if x < y {
			x = y
		}
		return x
	}
	res := 0
	for i, f := range flips {
		biggest = max(biggest, f)
		if biggest == i+1 {
			res++
		}
	}
	return res
}

func main() {
	flips := []int{4, 1, 2, 3}
	println(numTimesAllBlue(flips))
}
