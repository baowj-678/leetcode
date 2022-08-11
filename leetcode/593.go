/**
 * @author: baowj
 * @time: 2022/7/29 10:04:36
 */
package main

func validSquare(p1 []int, p2 []int, p3 []int, p4 []int) bool {
	edgeLen := 0
	diagLen := 0
	if distance(p1, p2) == distance(p1, p3) {
		edgeLen = distance(p1, p2)
		diagLen = distance(p1, p4)
	} else if distance(p1, p2) > distance(p1, p3) {
		diagLen = distance(p1, p2)
		edgeLen = distance(p1, p3)
		if distance(p1, p4) != edgeLen {
			return false
		}
		p2, p4 = p4, p2
	} else {
		diagLen = distance(p1, p3)
		edgeLen = distance(p1, p2)
		if distance(p1, p4) != edgeLen {
			return false
		}
		p3, p4 = p4, p3
	}
	if !(edgeLen == distance(p4, p2) && edgeLen == distance(p4, p3)) {
		return false
	}
	if !(diagLen == distance(p2, p3)) {
		return false
	}
	if !(edgeLen != 0 && edgeLen*2 == diagLen) {
		return false
	}
	return true
}

func distance(p1 []int, p2 []int) int {
	return (p1[0]-p2[0])*(p1[0]-p2[0]) + (p1[1]-p2[1])*(p1[1]-p2[1])
}
