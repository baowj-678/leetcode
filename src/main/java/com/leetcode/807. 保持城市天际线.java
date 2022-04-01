/**
 * @author BaoWJ
 * @date 2021/12/13
 * @link https://leetcode-cn.com/problems/max-increase-to-keep-city-skyline/
 */
package com.leetcode;

class Solution {
    public int maxIncreaseKeepingSkyline(int[][] grid) {
        int m = grid.length, n = grid[0].length;
        int[] mDirect = new int[m];
        int[] nDirect = new int[n];
        for (int i = 0; i < m; i++) {
            int maxHeight = 0;
            for (int j = 0; j < n; j++) {
                maxHeight = Math.max(maxHeight, grid[i][j]);
            }
            mDirect[i] = maxHeight;
        }
        for (int i = 0; i < n; i++) {
            int maxHeight = 0;
            for (int j = 0; j < m; j++) {
                maxHeight = Math.max(maxHeight, grid[j][i]);
            }
            nDirect[i] = maxHeight;
        }
        int ans = 0;
        for (int i = 0; i < m; i++) {
            for (int j = 0; j < n; j++) {
                ans += (Math.min(mDirect[i], nDirect[j]) - grid[i][j]);
            }
        }
        return ans;
    }
}