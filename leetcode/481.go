/**
 * @author: baowj
 * @time: 2022/11/1 17:05:19
 */
package main

func magicalString(n int) int {
	str := make([]byte, n)
	str[0] = 1
	if n > 1 {
		str[1] = 2
	}
	res := 0
	k := 0
	var flag byte = 1
	for i := 0; i < n; {
		if str[k] == 1 {
			str[i] = flag
			if str[i] == 1 {
				res++
			}
			i++
			k++
		} else {
			str[i] = flag
			if str[i] == 1 {
				res++
			}
			i++
			if i == n {
				break
			}
			str[i] = flag
			if str[i] == 1 {
				res++
			}
			i++
			k++
		}

		if flag == 1 {
			flag = 2
		} else {
			flag = 1
		}
	}
	return res
}
