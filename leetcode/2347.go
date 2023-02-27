/**
 * @author: baowj
 * @time: 2023/2/20 11:31
 */
package main

func bestHand(ranks []int, suits []byte) string {
	m := map[int]int{}
	for _, rank := range ranks {
		if _, ok := m[rank]; !ok {
			m[rank] = 0
		}
		m[rank]++
	}
	checkFlush := func() bool {
		for i := 1; i < len(suits); i++ {
			if suits[i] != suits[0] {
				return false
			}
		}
		return true
	}
	checkThreeOfAKind := func() bool {
		for _, v := range m {
			if v >= 3 {
				return true
			}
		}
		return false
	}

	checkPiar := func() bool {
		for _, v := range m {
			if v >= 2 {
				return true
			}
		}
		return false
	}

	if checkFlush() {
		return "Flush"
	} else if checkThreeOfAKind() {
		return "Three of a Kind"
	} else if checkPiar() {
		return "Pair"
	} else {
		return "High Card"
	}
}
