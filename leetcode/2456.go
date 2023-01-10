/**
 * @author: baowj
 * @time: 2022/11/2 17:05:25
 */
package main

import "sort"

type struct2456 struct {
	string
	int
}

func mostPopularCreator(creators []string, ids []string, views []int) [][]string {
	m := map[string]int{}
	maxId := map[string]string{}
	maxV := map[string]int{}
	for i := 0; i < len(creators); i++ {
		if v, ok := m[creators[i]]; ok {
			m[creators[i]] = v + views[i]
		} else {
			m[creators[i]] = views[i]
		}

		if v, ok := maxV[creators[i]]; ok {
			if v < views[i] || (v == views[i] && maxId[creators[i]] > ids[i]) {
				maxV[creators[i]] = views[i]
				maxId[creators[i]] = ids[i]
			}
		} else {
			maxV[creators[i]] = views[i]
			maxId[creators[i]] = ids[i]
		}
	}
	sortCreators := make([]struct2456, len(m))
	i := 0
	for k, v := range m {
		sortCreators[i].string = k
		sortCreators[i].int = v
		i++
	}
	sort.Slice(sortCreators, func(i, j int) bool {
		return sortCreators[i].int > sortCreators[j].int
	})
	maxView := sortCreators[0].int
	res := [][]string{}
	for i = 0; i < len(sortCreators); i++ {
		if sortCreators[i].int == maxView {
			res = append(res, []string{sortCreators[i].string, maxId[sortCreators[i].string]})
		} else {
			break
		}
	}
	return res
}
