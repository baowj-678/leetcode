/**
 * @author: baowj
 * @time: 2023/9/14 10:03
 */
package main

func queensAttacktheKing(queens [][]int, king []int) [][]int {
	const Queen, King, Attack = 1, 2, 3
	board := [8][8]int{}
	for _, q := range queens {
		board[q[0]][q[1]] = Queen
	}
	board[king[0]][king[1]] = King

	directs := [][2]int{{-1, 0}, {-1, 1}, {0, 1}, {1, 1}, {1, 0}, {1, -1}, {0, -1}, {-1, -1}}
	for _, direct := range directs {
		i, j := king[0]+direct[0], king[1]+direct[1]
		for 0 <= i && i < 8 && 0 <= j && j < 8 {
			if board[i][j] == Queen {
				board[i][j] = Attack
				break
			}
			i += direct[0]
			j += direct[1]
		}
	}
	var res [][]int
	for i := 0; i < 8; i++ {
		for j := 0; j < 8; j++ {
			if board[i][j] == Attack {
				res = append(res, []int{i, j})
			}
		}
	}
	return res
}
