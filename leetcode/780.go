/**
 * @author: baowj
 * @time: 2022/4/9 11:14:51
 */
package main

func reachingPoints(sx int, sy int, tx int, ty int) bool {
	for sx <= tx && sy <= ty {
		if sx == tx && sy == ty {
			return true
		}
		if tx > ty {
			tmp := (tx - sx) / ty
			if tmp > 1 {
				tx -= tmp * ty
			} else {
				tx -= ty
			}
		} else if tx < ty {
			tmp := (ty - sy) / tx
			if tmp > 1 {
				ty -= tmp * tx
			} else {
				ty -= tx
			}
		} else {
			return false
		}
	}
	return false
}

func main() {
	print(reachingPoints(1, 1, 8000, 78890))
}
