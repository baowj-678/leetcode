/**
 * @author: baowj
 * @time: 2022/7/30 10:08:32
 */
package main

func largestComponentSize(nums []int) int {
	parent := make([]int, len(nums))
	size := make([]int, len(nums))
	table := getLookUpTable(100000)
	m := map[int]int{}
	for i := 0; i < len(nums); i++ {
		parent[i] = i
		size[i] = 1
	}
	for i, x := range nums {
		for j := 0; j < len(table) && table[j] <= x; j++ {
			if x%table[j] == 0 {
				if index, ok := m[table[j]]; ok {
					merge(parent, size, i, index)
				}
				m[table[j]] = i
			}
		}
	}
	max := 0
	for _, x := range size {
		if x > max {
			max = x
		}
	}
	return max
}

func getLookUpTable(n int) []int {
	table := []int{2}
	for i := 3; i <= n; i++ {
		isPrime := true
		for _, x := range table {
			if i%x == 0 {
				isPrime = false
				break
			}
		}
		if isPrime {
			table = append(table, i)
		}
	}
	return table
}

func find(parent []int, x int) int {
	tmp := x
	for parent[tmp] != tmp {
		tmp = parent[tmp]
	}
	for parent[x] != x {
		x, parent[x] = parent[x], tmp
	}
	return tmp
}

func merge(parent []int, size []int, x int, y int) {
	yP := find(parent, y)
	if yP != x {
		parent[yP] = x
		size[x] += size[yP]
	}
}

func main() {
	nums := []int{4, 6, 15, 35}
	println(largestComponentSize(nums))
}
