/**
 * @author: baowj
 * @time: 2022/11/6 19:04:21
 */
package main

import "strings"

func interpret(command string) string {
	command = strings.Replace(command, "()", "o", -1)
	command = strings.Replace(command, "(al)", "al", -1)
	return command
}
