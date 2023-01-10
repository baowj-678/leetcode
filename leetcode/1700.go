/**
 * @author: baowj
 * @time: 2022/10/26 22:08:25
 */
package main

func countStudents(students []int, sandwiches []int) int {
	cnt := 0
	for len(students) > 0 && len(sandwiches) > 0 {
		if students[0] == sandwiches[0] {
			cnt = 0
			students = students[1:]
			sandwiches = sandwiches[1:]
		} else {
			cnt++
			if cnt >= len(students) {
				return len(students)
			}
			students = append(students[1:], students[0])
		}
	}
	return 0
}
