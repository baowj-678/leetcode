/**
 * @author: baowj
 * @time: 2023/9/27 10:44
 */
package main

import "sort"

func filterRestaurants(restaurants [][]int, veganFriendly int, maxPrice int, maxDistance int) []int {
	res := [][]int{}
	for _, r := range restaurants {
		if veganFriendly == 1 && r[2] == 0 {
			continue
		}
		if r[3] > maxPrice {
			continue
		}
		if r[4] > maxDistance {
			continue
		}
		res = append(res, r)
	}
	sort.Slice(res, func(i, j int) bool {
		if res[i][1] == res[j][1] {
			return res[i][0] > res[j][0]
		} else {
			return res[i][1] > res[j][1]
		}
	})

	ans := make([]int, 0, len(res))
	for i := range res {
		ans = append(ans, res[i][0])
	}
	return ans
}
