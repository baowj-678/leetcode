/**
 * @author: baowj
 * @time: 2022/7/14 10:51:17
 */
package main

type WordFilter struct {
	prefix *FixFilter
	suffix *FixFilter
}

type FixFilter struct {
	value []int
	table map[byte]*FixFilter
}

func Constructor745(words []string) WordFilter {
	prefix := new(FixFilter)
	prefix.table = make(map[byte]*FixFilter)

	suffix := new(FixFilter)
	suffix.table = make(map[byte]*FixFilter)

	for i := len(words) - 1; i >= 0; i-- {
		word := words[i]
		tmp := prefix
		tmp.value = append(tmp.value, i)
		for _, c := range word {
			byteC := byte(c)
			value, ok := tmp.table[byteC]
			if ok {
				tmp = value
				tmp.value = append(tmp.value, i)
			} else {
				tmp.table[byteC] = new(FixFilter)
				tmp = tmp.table[byteC]
				tmp.table = make(map[byte]*FixFilter)
				tmp.value = []int{i}
			}
		}

		tmp1 := suffix
		tmp1.value = append(tmp1.value, i)
		for j := len(word) - 1; j >= 0; j-- {
			c := word[j]
			value, ok := tmp1.table[c]
			if ok {
				tmp1 = value
				tmp1.value = append(tmp1.value, i)
			} else {
				tmp1.table[c] = new(FixFilter)
				tmp1 = tmp1.table[c]
				tmp1.table = make(map[byte]*FixFilter)
				tmp1.value = []int{i}
			}
		}
	}
	ans := WordFilter{
		prefix: prefix,
		suffix: suffix,
	}
	return ans
}

func (this *WordFilter) F(pref string, suff string) int {
	prefix := this.prefix
	for _, c := range pref {
		char := byte(c)
		tmp, ok := prefix.table[char]
		if ok {
			prefix = tmp
		} else {
			return -1
		}
	}
	suffix := this.suffix
	for i := len(suff) - 1; i >= 0; i-- {
		c := suff[i]
		tmp, ok := suffix.table[c]
		if ok {
			suffix = tmp
		} else {
			return -1
		}
	}

	i := 0
	j := 0
	for i < len(prefix.value) && j < len(suffix.value) {
		if prefix.value[i] == suffix.value[j] {
			return prefix.value[i]
		} else if prefix.value[i] > suffix.value[j] {
			i++
		} else {
			j++
		}
	}
	return -1
}

/**
 * Your WordFilter object will be instantiated and called as such:
 * obj := Constructor(words);
 * param_1 := obj.F(pref,suff);
 */

func main() {
	obj := Constructor745([]string{"apple"})
	println(obj.F("a", "e"))
}
