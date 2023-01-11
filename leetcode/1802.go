/**
 * @author: baowj
 * @time: 2023/1/11 17:26:12
 * @email: bwj_678@qq.com
 */
package main

func maxValue(n int, index int, maxSum int) int {
	res := 1
	maxSum -= n
	if maxSum == 0 {
		return res
	}
	//both
	left := index
	right := n - index - 1
	length := left
	if right < left {
		length = right
	}
	l, r := 0, length
	for {
		k := (l + r + 1) / 2
		sum := (1 + k*2 + 1) * (k + 1) / 2
		if maxSum > sum {
			l = k
		} else if maxSum < sum {
			r = k - 1
		} else {
			return res + k + 1
		}
		if l >= r {
			break
		}
	}
	maxSum -= (1 + r*2 + 1) * (r + 1) / 2
	res += r + 1
	if r < length {
		return res
	}
	length = r*2 + 1
	// single
	l, r = 0, AbsInt1802(left-right)
	for {
		k := (l + r + 1) / 2
		sum := (length + 1 + length + k) * k / 2
		if maxSum > sum {
			l = k
		} else if maxSum < sum {
			r = k - 1
		} else {
			return res + k
		}
		if l >= r {
			break
		}
	}
	maxSum -= (length + 1 + length + r) * r / 2
	res += r
	if r < AbsInt1802(left-right) {
		return res
	}
	// n
	res += maxSum / n
	return res
}

func AbsInt1802(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
