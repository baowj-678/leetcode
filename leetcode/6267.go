/**
 * @author: baowj
 * @time: 2022/12/18 16:46:16
 * @email: bwj_678@qq.com
 */
package main

import "strconv"

func isPossible(n int, edges [][]int) bool {
	hasEdge := map[string]bool{}
	vertexDegree := make([]int, n+1)
	for _, edge := range edges {
		if edge[0] > edge[1] {
			edge[0], edge[1] = edge[1], edge[0]
		}
		k := strconv.Itoa(edge[0]) + strconv.Itoa(edge[1])
		hasEdge[k] = true
		vertexDegree[edge[0]]++
		vertexDegree[edge[1]]++
	}
	oddVertex := []int{}
	for i, x := range vertexDegree {
		if x%2 == 1 {
			oddVertex = append(oddVertex, i)
		}
	}
	if len(oddVertex) > 4 {
		return false
	}
	if len(oddVertex) == 0 {
		return true
	}
	if len(oddVertex) == 2 {
		k := strconv.Itoa(oddVertex[0]) + strconv.Itoa(oddVertex[1])
		if _, ok := hasEdge[k]; !ok {
			return true
		}
		for v := 1; v <= n; v++ {
			if v == oddVertex[0] || v == oddVertex[1] {
				continue
			}
			if v < oddVertex[0] {
				k = strconv.Itoa(v) + strconv.Itoa(oddVertex[0])
			} else {
				k = strconv.Itoa(oddVertex[0]) + strconv.Itoa(v)
			}
			if _, ok := hasEdge[k]; ok {
				continue
			}

			if v < oddVertex[1] {
				k = strconv.Itoa(v) + strconv.Itoa(oddVertex[1])
			} else {
				k = strconv.Itoa(oddVertex[1]) + strconv.Itoa(v)
			}
			if _, ok := hasEdge[k]; !ok {
				return true
			}
		}
	}
	if len(oddVertex) == 4 {
		for i := 0; i < 4; i++ {
			for j := i + 1; j < 4; j++ {
				k := strconv.Itoa(oddVertex[i]) + strconv.Itoa(oddVertex[j])
				if _, ok := hasEdge[k]; ok {
					continue
				}
				for p := 0; p < 4; p++ {
					if p != i && p != j {
						for q := p + 1; q < 4; q++ {
							if q != i && q != j {
								k = strconv.Itoa(oddVertex[p]) + strconv.Itoa(oddVertex[q])
								if _, ok := hasEdge[k]; !ok {
									return true
								}
							}
						}
					}
				}
			}
		}
	}
	return false
}
