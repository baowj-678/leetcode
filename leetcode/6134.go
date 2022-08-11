/**
 * @author: baowj
 * @time: 2022/7/31 10:42:23
 */
package main

import "math"

func closestMeetingNode(edges []int, node1 int, node2 int) int {
	dis1 := make([]int, len(edges))
	dis2 := make([]int, len(edges))
	for i := 0; i < len(dis1); i++ {
		dis1[i] = math.MaxInt
		dis2[i] = math.MaxInt
	}
	dis1[node1] = 0
	dis2[node2] = 0
	queue := []int{}
	queue = append(queue, node1)
	for len(queue) > 0 {
		node := queue[0]
		queue = queue[1:]
		if edges[node] == -1 {
			continue
		} else {
			if dis1[node]+1 < dis1[edges[node]] {
				dis1[edges[node]] = dis1[node] + 1
				queue = append(queue, edges[node])
			}
		}
	}

	queue = []int{}
	queue = append(queue, node2)
	for len(queue) > 0 {
		node := queue[0]
		queue = queue[1:]
		if edges[node] == -1 {
			continue
		} else {
			if dis2[node]+1 < dis2[edges[node]] {
				dis2[edges[node]] = dis2[node] + 1
				queue = append(queue, edges[node])
			}
		}
	}

	min := math.MaxInt
	minNode := -1
	for i := 0; i < len(dis1); i++ {
		if dis1[i] < dis2[i] {
			dis1[i] = dis2[i]
		}
		if dis1[i] < min {
			min = dis1[i]
			minNode = i
		}
	}
	return minNode
}

func main() {
	nums := []int{5, 4, 5, 4, 3, 6, -1}
	println(closestMeetingNode(nums, 0, 1))
}
