/**
 * @author: baowj
 * @time: 2023/5/17 10:50
 */
package main

import "sort"

func rearrangeBarcodes(barcodes []int) []int {
	m := map[int]int{}
	for _, barcode := range barcodes {
		if _, ok := m[barcode]; !ok {
			m[barcode] = 0
		}
		m[barcode]++
	}
	distinctBarcodes := make([]int, len(m))
	i := 0
	for k, _ := range m {
		distinctBarcodes[i] = k
		i++
	}
	sort.Slice(distinctBarcodes, func(i, j int) bool {
		return m[distinctBarcodes[i]] > m[distinctBarcodes[j]]
	})

	i = 0
	for _, code := range distinctBarcodes {
		num := m[code]
		for j := 0; j < num; j++ {
			if i >= len(barcodes) {
				i = 1
			}
			barcodes[i] = code
			i += 2
		}
	}
	return barcodes
}
