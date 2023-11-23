package com.leetcode;

import java.util.Arrays;

class Solution2304 {
    public int minPathCost(int[][] grid, int[][] moveCost) {
        int len = grid[0].length;
        int[] dp = new int[len];
//        for (int i = 0; i < len; i++) {
//            dp[i] = Integer.MAX_VALUE;
//        }
        for (int t = 0; t < grid.length; t++) {
            int[] row = grid[t];
            int[] tmp = new int[len];
            for (int i = 0; i < len; i++) {
                tmp[i] = Integer.MAX_VALUE;
            }

            for (int i = 0; i < len; i++) {
                int v = row[i];
                if (t < grid.length - 1) {
                    for (int j = 0; j < len; j++) {
                        int cost = moveCost[v][j];
                        if (tmp[j] > dp[i] + v + cost) {
                            tmp[j] = dp[i] + v + cost;
                        }
                    }
                } else {
                    tmp[i] = dp[i] + v;
                }
            }

            for (int i = 0; i < len; i++) {
                dp[i] = tmp[i];
            }
        }
        return Arrays.stream(dp).min().getAsInt();
    }
}