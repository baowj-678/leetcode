package com.leetcode;

import javafx.util.Pair;

class Solution2661 {
    public int firstCompleteIndex(int[] arr, int[][] mat) {
        int m = mat.length;
        int n = mat[0].length;
        int[] mCnt = new int[m];
        int[] nCnt = new int[n];

        int[][] idx = new int[m * n][2];
        for (int i = 0; i < m; i++) {
            for (int j = 0; j < n; j++) {
                idx[mat[i][j]-1][0] = i;
                idx[mat[i][j]-1][1] = j;
            }
        }

        for (int i = 0; i < arr.length; i++) {
            int[] p = idx[arr[i]];
            mCnt[p[0]]++;
            if (mCnt[p[0]] == n) {
                return i;
            }
            nCnt[p[1]]++;
            if (nCnt[p[1]] == m) {
                return i;
            }
        }
        return -1;
    }
}