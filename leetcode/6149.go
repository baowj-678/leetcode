/**
 * @author: baowj
 * @time: 2022/8/14 10:41:32
 */
package main

func edgeScore(edges []int) int {
	score := make([]int, len(edges))
	for i := 0; i < len(score); i++ {
		score[i] = 0
	}
	for i := 0; i < len(edges); i++ {
		score[edges[i]] += i
	}
	max, maxId := 0, 0
	for i := 0; i < len(score); i++ {
		if score[i] > max {
			max = score[i]
			maxId = i
		}
	}
	return maxId
}
