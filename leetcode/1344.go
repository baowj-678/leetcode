/**
 * @author: baowj
 * @time: 2022/8/26 10:02:44
 */
package main

import "math"

func angleClock(hour int, minutes int) float64 {
	minuteAngle := float64(minutes * 6)
	var hourAngle float64 = 0
	hourAngle = float64(hour%12) * (360 / 12)
	hourAngle += (float64(minutes) / 60) * (360 / 12)
	if math.Abs(minuteAngle-hourAngle) <= 180 {
		return math.Abs(minuteAngle - hourAngle)
	} else {
		return 360 - math.Abs(minuteAngle-hourAngle)
	}
}

func main() {
	println(angleClock(12, 30))

}
