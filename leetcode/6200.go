/**
 * @author: baowj
 * @time: 2022/10/9 15:41:54
 */
package main

func hardestWorker(n int, logs [][]int) int {
	maxId, maxLength := 0, 0
	for i, lastTime := 0, 0; i < len(logs); i++ {
		log := logs[i]
		if log[1]-lastTime > maxLength || (log[1]-lastTime == maxLength && log[0] < maxId) {
			maxId = log[0]
			maxLength = log[1] - lastTime
		}
		lastTime = log[1]
	}
	return maxId
}
