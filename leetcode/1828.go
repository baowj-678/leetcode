/**
 * @author: baowj
 * @time: 2023/1/24 20:09:45
 * @email: bwj_678@qq.com
 */
package main

func countPoints(points [][]int, queries [][]int) []int {
	res := make([]int, len(queries))
	for i, querie := range queries {
		cnt := 0
		for _, point := range points {
			if (point[0]-querie[0])*(point[0]-querie[0])+(point[1]-querie[1])*(point[1]-querie[1]) <= querie[2]*querie[2] {
				cnt++
			}
		}
		res[i] = cnt
	}
	return res
}
