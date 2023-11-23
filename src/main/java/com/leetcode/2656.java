package com.leetcode;

import java.util.Arrays;

class Solution2656 {
    public int maximizeSum(int[] nums, int k) {
        int res = Arrays.stream(nums).max().getAsInt();
        return res * k + (k - 1) * k / 2;
    }
}