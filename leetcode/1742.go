/**
 * @author: baowj
 * @time: 2022/11/23 13:41
 */
package main

func countBalls(lowLimit int, highLimit int) int {
	m := map[int]int{}
	for i := lowLimit; i <= highLimit; i++ {
		sum := f1742(i)
		if v, ok := m[sum]; ok {
			m[sum] = v + 1
		} else {
			m[sum] = 1
		}
	}
	max := 0
	for _, v := range m {
		if max < v {
			max = v
		}
	}
	return max
}

func f1742(x int) int {
	res := 0
	for x > 0 {
		res += x % 10
		x /= 10
	}
	return res
}
