package com.leetcode;

import java.util.Arrays;

class Solution274 {
    public int hIndex(int[] citations) {
        int l = 0, r = 1000;
        while (l < r) {
            int mid = (l + r + 1) / 2;

            int cnt = 0;
            for (int c: citations) {
                if (c >= mid) {
                    cnt++;
                }
            }
            if (cnt >= mid) {
                l = mid;
            } else {
                r = mid - 1;
            }
        }
        return l;
    }
}