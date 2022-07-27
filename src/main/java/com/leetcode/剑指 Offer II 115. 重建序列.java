package com.leetcode;

import java.util.Arrays;

class SolutionOffer115 {
    public boolean sequenceReconstruction(int[] nums, int[][] sequences) {
        boolean[] flags = new boolean[nums.length+1];
        int[] indexs = new int[nums.length+1];
        for (int i = 0; i < nums.length; i++) {
            indexs[nums[i]] = i;
        }
        Arrays.fill(flags, false);
        for (int[] sequence : sequences) {
            for (int i = 0; i+1 < sequence.length; i++) {
                int index = indexs[sequence[i]];
                if (nums[index+1] == sequence[i+1]) {
                    flags[sequence[i]] = true;
                }
            }
        }
        int cnt = 0;
        for (int i = 1; i < nums.length; i++) {
            if (!flags[i]) {
                cnt++;
            }
        }
        return cnt == 1;
    }
}