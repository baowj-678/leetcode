/**
 * @author: baowj
 * @time: 2023/6/21 10:11
 */
package main

func flipChess(chessboard []string) int {
	h, w := len(chessboard), len(chessboard[0])
	byteBoard := make([][]int, h)
	for i := 0; i < h; i++ {
		byteBoard[i] = make([]int, w)
	}

	for i := 0; i < h; i++ {
		for j := 0; j < w; j++ {
			switch chessboard[i][j] {
			case 'X':
				byteBoard[i][j] = 1
			case 'O':
				byteBoard[i][j] = -1
			case '.':
				byteBoard[i][j] = 0
			}
		}
	}

	check := func(x, y int, board [][]int) (bool, [][2]int) {
		directions := [][2]int{{1, 0}, {-1, 0}, {0, 1}, {0, -1}, {1, 1}, {1, -1}, {-1, 1}, {-1, -1}}
		board[x][y] = 1
		q := [][2]int{}
		for _, direction := range directions {
			tmpQ := [][2]int{}
			isOk := false
			tmpX, tmpY := x+direction[0], y+direction[1]
			for 0 <= tmpX && tmpX < h && 0 <= tmpY && tmpY < w {
				if board[tmpX][tmpY] == 0 {
					isOk = false
					break
				} else if board[tmpX][tmpY] == -1 {
					tmpQ = append(tmpQ, [2]int{tmpX, tmpY})
				} else {
					isOk = true
					break
				}
				tmpX, tmpY = tmpX+direction[0], tmpY+direction[1]
			}
			if isOk {
				for _, p := range tmpQ {
					board[p[0]][p[1]] = 1
				}
				q = append(q, tmpQ...)
			}
		}
		return len(q) > 0, q
	}

	newBoard := func() [][]int {
		board := make([][]int, h)
		for i := 0; i < h; i++ {
			board[i] = make([]int, w)
			copy(board[i], byteBoard[i])
		}
		return board
	}

	max := func(x, y int) int {
		if x < y {
			x = y
		}
		return x
	}

	res := 0
	// try everywhere
	for i := 0; i < h; i++ {
		for j := 0; j < w; j++ {
			if byteBoard[i][j] == 0 {
				board := newBoard()
				q := [][2]int{{i, j}}
				tmp := 0
				for len(q) > 0 {
					point := q[0]
					q = q[1:]
					if ok, appendQ := check(point[0], point[1], board); ok {
						tmp += len(appendQ)
						q = append(q, appendQ...)
					}
				}
				res = max(res, tmp)
			}
		}
	}
	return res
}
