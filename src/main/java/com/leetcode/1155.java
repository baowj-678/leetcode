package com.leetcode;

class Solution1155 {
    public int numRollsToTarget(int n, int k, int target) {
        long MAX = 1000000007;
        long[] dp = new long[target + 1];
        for (int i = 1; i <= k && i < dp.length; i++) {
            dp[i] = 1;
        }
        for (int i = 1; i < n; i++) {
            for (int j = dp.length-1; j > 0; j--) {
                dp[j] = 0;
                for (int t = 1; t <= k; t++) {
                    if (j - t <= 0) {
                        break;
                    }
                    dp[j] = (dp[j - t] + dp[j]) % MAX;
                }
            }
        }
        return (int)dp[target];
    }

    public static void main(String[] args) {
        Solution1155 so = new Solution1155();
        System.out.println(so.numRollsToTarget(4, 6, 7));
    }
}