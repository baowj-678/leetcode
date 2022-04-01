/**
 * @author Bao WJ
 * @link https://leetcode-cn.com/problems/longest-harmonious-subsequence/
 * @date 2021/11/20
 */
package com.leetcode;

import java.util.Arrays;

class Solution594 {
    public int findLHS(int[] nums) {
        Arrays.sort(nums);
        int res = 0;
        for (int i = 0, j = 0; j < nums.length; j++) {
            while (i < j && nums[j] - nums[i] > 1) {
                i++;
            }
            if (nums[j] - nums[i] == 1) {
                res = Math.max(j - i + 1, res);
            }
        }
        return res;
    }
}