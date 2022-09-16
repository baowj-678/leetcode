/**
 * @author: baowj
 * @time: 2022/8/28 21:49:33
 */
package main

func garbageCollection(garbage []string, travel []int) int {
	res := 0
	mp, pp, gp := 0, 0, 0
	for i, s := range garbage {
		m_, p_, g_ := getMPG(s)
		if m_ > 0 {
			for j := mp; j < i; j++ {
				res += travel[j]
			}
			mp = i
			res += m_
		}
		if p_ > 0 {
			for j := pp; j < i; j++ {
				res += travel[j]
			}
			pp = i
			res += p_
		}
		for g_ > 0 {
			for j := gp; j < i; j++ {
				res += travel[j]
			}
			gp = i
			res += g_
		}
	}
	return res
}

func getMPG(s string) (int, int, int) {
	m, p, g := 0, 0, 0
	for _, c := range s {
		switch c {
		case 'M':
			m++
		case 'P':
			p++
		case 'G':
			g++
		}
	}
	return m, p, g
}
