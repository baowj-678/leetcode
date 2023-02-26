/**
 * @author: baowj
 * @time: 2023/2/18 13:10:05
 * @email: bwj_678@qq.com
 */
package main

import "sort"

/**
 * This is the declaration of customFunction API.
 * @param  x    int
 * @param  x    int
 * @return 	    Returns f(x, y) for any given positive integers x and y.
 *			    Note that f(x, y) is increasing with respect to both x and y.
 *              i.e. f(x, y) < f(x + 1, y), f(x, y) < f(x, y + 1)
 */

func findSolution(customFunction func(int, int) int, z int) [][]int {
	res := [][]int{}
	for x := 1; x <= 1000; x++ {
		y := sort.Search(1000, func(i int) bool {
			return customFunction(x, i+1) >= z
		})
		y++
		if customFunction(x, y) == z {
			res = append(res, []int{x, y})
		}
	}
	return res
}
