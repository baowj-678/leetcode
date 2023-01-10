/**
 * @author: baowj
 * @time: 2023/1/9 18:33:58
 * @email: bwj_678@qq.com
 */
package main

func reinitializePermutation(n int) int {
	n--
	perm := make([]int, n)
	for i := 0; i < n; i++ {
		perm[i] = i
	}
	res := 0
	array := make([]int, n)
	for {
		res++
		isOk := true
		for i := 0; i < n; i++ {
			array[(2*i)%n] = perm[i]
			if array[(2*i)%n] != (2*i)%n {
				isOk = false
			}
		}
		array, perm = perm, array
		if isOk {
			return res
		}
	}
}

func main() {
	println(reinitializePermutation(4))
}
