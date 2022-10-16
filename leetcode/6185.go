/**
 * @author: baowj
 * @time: 2022/9/18 10:29:55
 */
package main

import "sort"

func matchPlayersAndTrainers(players []int, trainers []int) int {
	sort.Ints(players)
	sort.Ints(trainers)
	res := 0
	k := 0
	for i := 0; i < len(players) && k < len(trainers); i++ {
		for ; k < len(trainers); k++ {
			if trainers[k] >= players[i] {
				res++
				k++
				break
			}
		}
	}
	return res
}
