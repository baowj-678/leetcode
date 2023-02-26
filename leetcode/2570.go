/**
 * @author: baowj
 * @time: 2023/2/25 19:54:23
 * @email: bwj_678@qq.com
 */
package main

func mergeArrays(nums1 [][]int, nums2 [][]int) [][]int {
	res := [][]int{}
	for i, j := 0, 0; i <= len(nums1) && j <= len(nums2); {
		if i == len(nums1) && j == len(nums2) {
			break
		} else if j == len(nums2) {
			res = append(res, nums1[i])
			i++
			continue
		} else if i == len(nums1) {
			res = append(res, nums2[j])
			j++
			continue
		}
		if nums1[i][0] == nums2[j][0] {
			res = append(res, []int{nums1[i][0], nums1[i][1] + nums2[j][1]})
			i++
			j++
		} else if nums1[i][0] < nums2[j][0] {
			res = append(res, nums1[i])
			i++
		} else {
			res = append(res, nums2[j])
			j++
		}
	}
	return res
}
