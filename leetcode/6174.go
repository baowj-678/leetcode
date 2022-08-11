/**
 * @author: baowj
 * @time: 2022/8/7 11:07:36
 */
package main

func taskSchedulerII(tasks []int, space int) int64 {
	var timeNow int64 = 0
	m := map[int]int64{}
	for _, task := range tasks {
		timeNow += 1
		if prevTime, ok := m[task]; ok {
			if timeNow-prevTime <= int64(space) {
				timeNow = int64(space) + prevTime + 1
			}
			m[task] = timeNow
		} else {
			m[task] = timeNow
		}
	}
	return timeNow
}

func main() {
	tasks := []int{1, 2, 1, 2, 3, 1}
	println(taskSchedulerII(tasks, 3))
}
