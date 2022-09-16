/**
 * @author: baowj
 * @time: 2022/9/15 10:34:57
 */
package main

func flipLights(n int, presses int) int {
	switch presses {
	case 0:
		return 1
	case 1:
		switch n {
		case 1:
			return 2
		case 2:
			return 3
		default:
			return 4
		}
	case 2:
		switch n {
		case 1:
			return 2
		case 2:
			return 4
		default:
			return 7
		}
	default:
		switch n {
		case 1:
			return 2
		case 2:
			return 4
		default:
			return 8
		}
	}
	return 8
}
