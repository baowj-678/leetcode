package com.leetcode;

class Solution2216 {
    public int minDeletion(int[] nums) {
        int last = -1, res = 0;
        boolean isEven = true;
        for (int num: nums) {
            if (isEven) {
                last = num;
                isEven = false;
            } else {
                if (last == num) {
                    res++;
                } else {
                    isEven = true;
                }
            }
        }
        if (!isEven) {
            res++;
        }
        return res;
    }
}