/**
 * @author: baowj
 * @time: 2022/8/31 21:40:09
 */
package main

func validateStackSequences(pushed []int, popped []int) bool {
	q := []int{}
	i := 0
	for _, x := range pushed {
		q = append(q, x)
		for len(q) > 0 && i < len(popped) && popped[i] == q[len(q)-1] {
			i++
			q = q[:len(q)-1]
		}
	}
	if len(q) == 0 && i == len(popped) {
		return true
	}
	return false
}
