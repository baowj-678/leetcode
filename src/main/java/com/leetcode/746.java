package com.leetcode;

class Solution746 {
    public int minCostClimbingStairs(int[] cost) {
        int[] dp = new int[cost.length + 1];
        for (int i = 2; i < dp.length; i++) {
            int v = dp[i-1] + cost[i-1];
            if (i - 2 >= 0) {
                v = Math.min(v, dp[i-2] + cost[i-2]);
            }
            dp[i] = v;
        }
        return dp[dp.length-1];
    }
}