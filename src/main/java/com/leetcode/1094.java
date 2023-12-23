package com.leetcode;

class Solution1094 {
    public boolean carPooling(int[][] trips, int capacity) {
        int[] passengers = new int[1001];
        for (int[] trip: trips) {
            passengers[trip[1]] += trip[0];
            passengers[trip[2]] -= trip[0];
        }
        int cnt = 0;
        for (int i: passengers) {
            cnt += i;
            if (cnt > capacity) {
                return false;
            }
        }
        return true;
    }
}