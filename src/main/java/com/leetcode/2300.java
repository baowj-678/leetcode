package com.leetcode;

import java.util.Arrays;

class Solution2300 {
    public int[] successfulPairs(int[] spells, int[] potions, long success) {
        Arrays.sort(potions);
        System.out.println(Arrays.toString(potions));
        for (int i = 0; i < spells.length; i++) {
            long v = (long) spells[i];
            int l = 0, r = potions.length;
            while (l < r) {
                int mid = (l + r) / 2;
                if (v * potions[mid] < success) {
                    l = mid + 1;
                } else {
                    r = mid;
                }
            }
            spells[i] = potions.length - l;
        }
        return spells;
    }
}