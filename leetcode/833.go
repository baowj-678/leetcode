package main

import (
	"sort"
	"strings"
)

type struct833 struct {
	index  int
	source string
	target string
}

type sortStruct833 []*struct833

func (s *sortStruct833) Len() int {
	return len(*s)
}

func (s *sortStruct833) Less(i int, j int) bool {
	return (*s)[i].index < (*s)[j].index
}

func (s *sortStruct833) Swap(i int, j int) {
	(*s)[i], (*s)[j] = (*s)[j], (*s)[i]
}

func findReplaceString(s string, indices []int, sources []string, targets []string) string {
	list := sortStruct833(make([]*struct833, len(indices)))
	for i := 0; i < len(indices); i++ {
		list[i] = &struct833{
			index:  indices[i],
			source: sources[i],
			target: targets[i],
		}
	}
	sort.Sort(&list)

	check := func(l int, subs string) bool {
		if l+len(subs) > len(s) {
			return false
		}
		for i := 0; i < len(subs); i++ {
			if s[l+i] != subs[i] {
				return false
			}
		}
		return true
	}

	res := strings.Builder{}
	i := 0
	for _, elem := range list {
		if elem.index > i {
			res.WriteString(s[i:elem.index])
			i = elem.index
		}
		if check(elem.index, elem.source) {
			res.WriteString(elem.target)
			i = i + len(elem.source)
		}
	}
	if i < len(s) {
		res.WriteString(s[i:])
	}
	return res.String()
}
