package com.leetcode;

import java.util.ArrayList;
import java.util.Arrays;
import java.util.Collections;
import java.util.List;

class Solution2578 {
    public int splitNum(int num) {
        ArrayList<Integer> nums = new ArrayList<>();
        while (num > 0) {
            nums.add(num % 10);
            num /= 10;
        }
        Collections.sort(nums);
        int v1 = 0, v2 = 0;
        for (int k = 0; k < nums.size(); k++) {
            if (k % 2 == 0) {
                v1 = v1 * 10 + nums.get(k);
            } else {
                v2 = v2 * 10 + nums.get(k);
            }
        }
        return v1 + v2;
    }

    public static void main(String[] args) {
        Solution2578 so = new Solution2578();
        System.out.println(so.splitNum(4325));
    }
}