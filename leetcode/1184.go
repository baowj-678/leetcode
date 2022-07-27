/**
 * @author: baowj
 * @time: 2022/7/24 09:48:16
 */
package main

func distanceBetweenBusStops(distance []int, start int, destination int) int {
	if start > destination {
		start, destination = destination, start
	}
	sum := 0
	dis := 0
	for i := 0; i < len(distance); i++ {
		sum += distance[i]
		if start <= i && i < destination {
			dis += distance[i]
		}
	}
	if sum-dis > dis {
		return sum - dis
	} else {
		return dis
	}
}
