package com.leetcode;

class Solution2609 {
    public int findTheLongestBalancedSubstring(String s) {
        int res = 0;
        int[] cnt = new int[2];
        for (char c: s.toCharArray()) {
            switch (c) {
                case '0':
                    if (cnt[1] > 0) {
                        cnt = new int[2];
                    }
                    cnt[0]++;
                    break;
                case '1':
                    cnt[1]++;
                    if (cnt[0] >= cnt[1]) {
                        res = Math.max(res, cnt[1] * 2);
                    } else {
                        cnt = new int[2];
                    }
                    break;
            }
        }
        return res;
    }
}