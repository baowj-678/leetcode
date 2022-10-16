/**
 * @author: baowj
 * @time: 2022/9/17 20:59:24
 */
package main

import (
	"fmt"
	"math"
	"sort"
)

func mostBooked(n int, meetings [][]int) int {
	rooms := make([]int, n)
	cnt := make([]int, n)
	sort.Slice(meetings, func(i, j int) bool {
		return meetings[i][0] < meetings[j][0]
	})
	for _, meeting := range meetings {
		earliestRoom := 0
		earliestTime := math.MaxInt
		hasRoom := false
		for i := 0; i < n; i++ {
			if meeting[0] >= rooms[i] {
				rooms[i] = meeting[1]
				cnt[i]++
				hasRoom = true
				break
			} else {
				if earliestTime > rooms[i] {
					earliestTime = rooms[i]
					earliestRoom = i
				}
			}
		}
		if !hasRoom {
			rooms[earliestRoom] += meeting[1] - meeting[0]
			cnt[earliestRoom]++
		}
	}
	fmt.Println(cnt)
	maxRoom, maxTimes := 0, 0
	for i := 0; i < n; i++ {
		if cnt[i] > maxTimes {
			maxTimes = cnt[i]
			maxRoom = i
		}
	}
	return maxRoom
}
