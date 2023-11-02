package com.leetcode;

class Solution2698 {
    public int punishmentNumber(int n) {
        int sum = 0;
        for (int i = 1; i <= n; i++) {
            int v = i * i;
            if (bfs(v, i)) {
                sum += v;
            }
        }
        return sum;
    }

    private boolean bfs (int v, int target) {
        for (int i = 10; v / i > 0; i *= 10) {
            int r = v % i;
            int l = v / i;
            if (target - r > 0 && bfs(l, target - r)) {
                return true;
            }
        }
        return v == target;
    }
}