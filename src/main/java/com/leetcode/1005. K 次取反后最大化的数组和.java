/**
 * @author Bao WJ
 * @date 2021/12/3
 * @link https://leetcode-cn.com/problems/maximize-sum-of-array-after-k-negations/
 */
package com.leetcode;

import java.util.Arrays;

class Solution1005 {
    public int largestSumAfterKNegations(int[] nums, int k) {
        Arrays.sort(nums);
        System.out.println(nums);
        int sum = 0, min = Integer.MAX_VALUE;
        for (int num : nums) {
            if (num >= 0) {
                sum += num;
            } else {
                if (k > 0) {
                    sum -= num;
                    k--;
                } else {
                    sum += num;
                }
            }
            min = Math.min(min, Math.abs(num));
        }
        if (k % 2 == 0) {
            return sum;
        } else {
            return sum - min;
        }
    }
}