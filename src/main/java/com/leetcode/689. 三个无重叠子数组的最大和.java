/**
 * @author BaoWJ
 * @date 2021/12/8
 * @link https://leetcode-cn.com/problems/maximum-sum-of-3-non-overlapping-subarrays/
 */
package com.leetcode;

import java.lang.reflect.Array;
import java.util.Arrays;

class Solution689 {
    public int[] maxSumOfThreeSubarrays(int[] nums, int k) {
        int sum = 0;
        for (int i = 0; i < nums.length; i++) {
            sum += nums[i];
            if (i >= k - 1) {
                int tmp = nums[i - k + 1];
                nums[i - k + 1] = sum;
                sum -= tmp;
            }
        }
        int[] ans = new int[3];
        int[] dp = new int[nums.length];
        int[][] dpIndex = new int[nums.length][2];
        int maxIndex = nums.length - k;
        Arrays.fill(dp, 0);
        ans[0] = ans[1] = ans[2] = 0;
        int maxV = 0;
        for (int i = nums.length - 2 * k; i >= 0; i--) {
            if (maxV <= nums[i + k]) {
                maxV = nums[i + k];
                maxIndex = i + k;
            }
            if (maxV + nums[i] >= dp[i + 1]) {
                dp[i] = maxV + nums[i];
                dpIndex[i][0] = i;
                dpIndex[i][1] = maxIndex;
            } else {
                dp[i] = dp[i + 1];
                dpIndex[i] = dpIndex[i + 1];
            }
        }
        maxV = 0;
        for (int i = 0; i <= nums.length - 3 * k; i++) {
            if (nums[i] + dp[i + k] > maxV) {
                maxV = nums[i] + dp[i + k];
                maxIndex = i;
            }
        }
        ans[0] = maxIndex;
        ans[1] = dpIndex[maxIndex + k][0];
        ans[2] = dpIndex[maxIndex + k][1];
        return ans;
    }
}