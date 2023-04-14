/**
 * @author: baowj
 * @time: 2023/4/11 9:58
 */
package main

func isRobotBounded(instructions string) bool {
	direction, x, y := 0, 0, 0
	for _, i := range instructions {
		switch i {
		case 'G':
			switch direction {
			case 0:
				y++
			case 1:
				x--
			case 2:
				y--
			case 3:
				x++
			}
		case 'L':
			direction = (direction + 1) % 4
		case 'R':
			direction = (direction - 1) % 4
		}
	}
	if x != 0 || y != 0 {
		return direction != 0
	} else {
		return true
	}
}
