package main

func isWinner(player1 []int, player2 []int) int {
	score1, score2 := 0, 0
	lastIdx := -10
	for i := 0; i < len(player1); i++ {
		if i-lastIdx <= 2 {
			score1 += player1[i]
		}
		score1 += player1[i]
		if player1[i] == 10 {
			lastIdx = i
		}
	}

	lastIdx = -10
	for i := 0; i < len(player2); i++ {
		if i-lastIdx <= 2 {
			score2 += player2[i]
		}
		score2 += player2[i]
		if player2[i] == 10 {
			lastIdx = i
		}
	}

	if score1 > score2 {
		return 1
	} else if score2 > score1 {
		return 2
	}
	return 0
}
