/**
 * @author: baowj
 * @time: 2022/10/26 22:25:17
 */
package main

func totalFruit(fruits []int) int {
	m := map[int]int{}
	res := 0
	for i, j := 0, 0; i < len(fruits) && j < len(fruits); j++ {
		if v, ok := m[fruits[j]]; ok {
			m[fruits[j]] = v + 1
		} else {
			m[fruits[j]] = 1
		}
		for len(m) > 2 {
			m[fruits[i]]--
			if m[fruits[i]] == 0 {
				delete(m, fruits[i])
			}
			i++
		}
		if j-i+1 > res {
			res = j - i + 1
		}
	}
	return res
}
