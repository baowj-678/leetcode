package com.leetcode;

class Solution {
    public int countDigits(int num) {
        int copy = num;
        int res = 0;
        while (copy > 0) {
            int tmp = copy % 10;
            copy /= 10;
            if (num % tmp == 0) {
                res++;
            }
        }
        return res;
    }
}