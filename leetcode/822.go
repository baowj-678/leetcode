/**
 * @author: baowj
 * @time: 2023/8/2 10:06
 */
package main

import (
	"fmt"
	"sort"
)

func flipgame(fronts []int, backs []int) int {
	exclude := map[int]struct{}{}
	for i := 0; i < len(fronts); i++ {
		if fronts[i] == backs[i] {
			exclude[fronts[i]] = struct{}{}
		}
	}
	sort.Ints(fronts)
	sort.Ints(backs)
	i, j := 0, 0
	for i < len(fronts) && j < len(backs) {
		if fronts[i] < backs[j] {
			if _, ok := exclude[fronts[i]]; !ok {
				return fronts[i]
			}
			i++
		} else {
			if _, ok := exclude[backs[j]]; !ok {
				return backs[j]
			}
			j++
		}
	}
	for i < len(fronts) {
		if _, ok := exclude[fronts[i]]; !ok {
			return fronts[i]
		}
		i++
	}

	for j < len(backs) {
		if _, ok := exclude[backs[j]]; !ok {
			return backs[j]
		}
		j++
	}

	return 0
}

func main() {
	fronts := []int{1}
	backs := []int{2}
	fmt.Println(flipgame(fronts, backs))
}
