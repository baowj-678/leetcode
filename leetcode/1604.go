/**
 * @author: baowj
 * @time: 2023/2/7 21:17:55
 * @email: bwj_678@qq.com
 */
package main

import (
	"sort"
	"strconv"
)

func alertNames(keyName []string, keyTime []string) []string {
	res := []string{}
	data := make([][2]string, len(keyName))
	for i := 0; i < len(keyName); i++ {
		data[i][0] = keyName[i]
		data[i][1] = keyTime[i]
	}

	sort.Slice(data, func(i, j int) bool {
		if data[i][0] == data[j][0] {
			return data[i][1] < data[j][0]
		}
		return data[i][0] < data[j][0]
	})

	isWithInAHour := func(time1, time2 string) bool {
		th1, _ := strconv.Atoi(time1[:2])
		tm1, _ := strconv.Atoi(time1[3:])
		th2, _ := strconv.Atoi(time2[:2])
		tm2, _ := strconv.Atoi(time2[3:])
		if th1 > th2 {
			return false
		} else if th1 == th2 {
			return true
		} else if th1+1 < th2 {
			return false
		} else {
			if tm2 > tm1 {
				return false
			} else {
				return true
			}
		}
	}
	for i := 0; i < len(data); i++ {
		name, time := data[i][0], data[i][1]
		if i > 1 {
			if data[i-2][0] == name {
				if isWithInAHour(data[i-2][1], time) {
					if len(res) == 0 || res[len(res)-1] != name {
						res = append(res, name)
					}
				}
			}
		}
	}
	sort.Strings(res)
	return res
}
