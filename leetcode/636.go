/**
 * @author: baowj
 * @time: 2022/8/7 10:00:24
 */
package main

import (
	"fmt"
	"strconv"
	"strings"
)

type Entry struct {
	functionId int
	action     string
	timestamp  int
}

func exclusiveTime(n int, logs []string) []int {
	res := make([]int, n)
	for i := 0; i < len(res); i++ {
		res[i] = 0
	}
	stack := []Entry{}
	timeNow := -1
	for _, log := range logs {
		entry := parseLog(log)
		if entry.action == "end" {
			stack = stack[:len(stack)-1]
			res[entry.functionId] += entry.timestamp - timeNow
			timeNow = entry.timestamp
		} else {
			start := Entry{
				timestamp: timeNow + 1,
			}
			if len(stack) > 0 {
				start.functionId = stack[len(stack)-1].functionId
			}
			stack = append(stack, entry)
			res[start.functionId] += entry.timestamp - start.timestamp
			timeNow = entry.timestamp - 1
		}
	}
	return res
}

func parseLog(str string) Entry {
	res := Entry{}
	strArray := strings.Split(str, ":")
	res.functionId, _ = strconv.Atoi(strArray[0])
	res.action = strArray[1]
	res.timestamp, _ = strconv.Atoi(strArray[2])
	return res
}

func main() {
	logs := []string{"0:start:0", "0:start:2", "0:end:5", "0:start:6", "0:end:6", "0:end:7"}
	fmt.Println(exclusiveTime(1, logs))
}
