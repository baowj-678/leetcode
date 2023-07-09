/**
 * @author: baowj
 * @time: 2023/6/25 14:11
 */
package main

import "fmt"

func checkOverlap(radius int, xCenter int, yCenter int, x1 int, y1 int, x2 int, y2 int) bool {
	checkLine := func(x1, y1, x2, y2 int) bool {
		if (xCenter-x1)*(xCenter-x1)+(yCenter-y1)*(yCenter-y1) <= radius*radius {
			return true
		}
		if (xCenter-x2)*(xCenter-x2)+(yCenter-y2)*(yCenter-y2) <= radius*radius {
			return true
		}
		if (x1 < xCenter && xCenter < x2) || (x2 < xCenter && xCenter < x1) {
			if (yCenter-y1)*(yCenter-y1) <= radius*radius {
				return true
			}
		}
		if (y1 < yCenter && yCenter < y2) || (y2 < yCenter && yCenter < y1) {
			if (xCenter-x1)*(xCenter-x1) <= radius*radius {
				return true
			}
		}
		return false
	}
	if checkLine(x2, y2, x2, y1) {
		return true
	}
	if checkLine(x1, y1, x2, y1) || checkLine(x1, y1, x1, y2) || checkLine(x1, y2, x2, y2) || checkLine(x2, y2, x2, y1) {
		return true
	}
	if y1 <= yCenter && yCenter <= y2 && x1 <= xCenter && xCenter <= x2 {
		return true
	}
	return false
}

func main() {
	fmt.Println(checkOverlap(4, 102, 50, 0, 0, 100, 100))
}
