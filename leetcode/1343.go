/**
 * @author: baowj
 * @time: 2022/8/26 10:21:10
 */
package main

func numOfSubarrays(arr []int, k int, threshold int) int {
	sum := 0
	res := 0
	for i, x := range arr {
		sum += x
		if i-k >= 0 {
			sum -= arr[i-k]
		}
		if i >= k-1 && sum/k >= threshold {
			res++
		}
	}
	return res
}
