/**
 * @author: baowj
 * @time: 2023/7/19 10:08
 */
package main

import (
	"sort"
)

func robotSim(commands []int, obstacles [][]int) int {
	x, y, p := 0, 0, 0
	sort.Slice(obstacles, func(i, j int) bool {
		if obstacles[i][0] < obstacles[j][0] {
			return true
		} else if obstacles[i][0] > obstacles[j][0] {
			return false
		} else {
			return obstacles[i][1] <= obstacles[j][1]
		}
	})

	xobstacles := map[int][]int{}
	yobstacles := map[int][]int{}
	for _, obstacle := range obstacles {
		tmpx, tmpy := obstacle[0], obstacle[1]
		if v, ok := xobstacles[tmpx]; ok {
			xobstacles[tmpx] = append(v, tmpy)
		} else {
			xobstacles[tmpx] = []int{tmpy}
		}

		if v, ok := yobstacles[tmpy]; ok {
			yobstacles[tmpy] = append(v, tmpx)
		} else {
			yobstacles[tmpy] = []int{tmpx}
		}
	}
	res := 0

	update := func() {
		if res < x*x+y*y {
			res = x*x + y*y
		}
	}
	for _, command := range commands {
		if command == -1 {
			p = (p + 1) % 4
		} else if command == -2 {
			p = (p + 3) % 4
		} else {
			switch p {
			case 0:
				if v, ok := xobstacles[x]; ok {
					oy := sort.Search(len(v), func(i int) bool {
						return v[i] > y
					})
					if oy == len(v) {
						y += command
					} else {
						if v[oy] > y+command {
							y += command
						} else {
							y = v[oy] - 1
						}
					}
				} else {
					y += command
				}
			case 1:
				if v, ok := yobstacles[y]; ok {
					ox := sort.Search(len(v), func(i int) bool {
						return v[i] > x
					})
					if ox == len(v) {
						x += command
					} else {
						if v[ox] > x+command {
							x += command
						} else {
							x = v[ox] - 1
						}
					}
				} else {
					x += command
				}
			case 2:
				if v, ok := xobstacles[x]; ok {
					oy := sort.Search(len(v), func(i int) bool {
						return v[i] >= y
					})
					if oy == 0 {
						y -= command
					} else {
						oy--
						if v[oy] < y-command {
							y -= command
						} else {
							y = v[oy] + 1
						}
					}
				} else {
					y -= command
				}
			case 3:
				if v, ok := yobstacles[y]; ok {
					ox := sort.Search(len(v), func(i int) bool {
						return v[i] >= x
					})
					if ox == 0 {
						x -= command
					} else {
						ox--
						if v[ox] < x-command {
							x -= command
						} else {
							x = v[ox] + 1
						}
					}
				} else {
					x -= command
				}
			}
			update()
		}
	}
	return res
}
