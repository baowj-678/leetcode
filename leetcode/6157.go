/**
 * @author: baowj
 * @time: 2022/8/21 10:36:19
 */
package main

func secondsToRemoveOccurrences(s string) int {
	zeroCnt, totalTime, prevIdx := 0, 0, -1
	for i, c := range s {
		if c == '1' {
			zeroTmp := 0
			for j := prevIdx + 1; j < i; j++ {
				if s[j] == '0' {
					zeroTmp++
				}
			}
			if zeroTmp >= totalTime {
				totalTime = zeroTmp + zeroCnt
			} else {
				if zeroTmp+zeroCnt <= totalTime {
					totalTime++
				} else {
					totalTime = zeroTmp + zeroCnt
				}
			}
			zeroCnt += zeroTmp
			prevIdx = i
		}
	}
	return totalTime
}

func main() {
	print(secondsToRemoveOccurrences("0110101"))
}
