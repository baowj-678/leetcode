package main

func minLength(s string) int {
	res := []byte(s)
	for {
		i, j := 0, 0
		for i < len(res) && j < len(res) {
			if i+1 < len(res) {
				if (res[i] == 'A' && res[i+1] == 'B') || (res[i] == 'C' && res[i+1] == 'D') {
					i += 2
					continue
				}
			}
			res[j] = res[i]
			i++
			j++
		}
		if i == j {
			break
		}
		res = res[:j]
	}
	return len(res)
}
