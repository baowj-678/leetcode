package com.leetcode;

import java.util.Arrays;

class Solution1465 {
    public int maxArea(int h, int w, int[] horizontalCuts, int[] verticalCuts) {
        Arrays.sort(horizontalCuts);
        int maxH = horizontalCuts[0];
        for (int i = 1; i < horizontalCuts.length; i++) {
            maxH = Math.max(maxH, horizontalCuts[i] - horizontalCuts[i-1]);
        }
        maxH = Math.max(maxH, h - horizontalCuts[horizontalCuts.length-1]);

        Arrays.sort(verticalCuts);
        int maxV = verticalCuts[0];
        for (int i = 1; i < verticalCuts.length; i++) {
            maxV = Math.max(maxV, verticalCuts[i] - verticalCuts[i-1]);
        }
        maxV = Math.max(maxV, w - verticalCuts[verticalCuts.length-1]);

        return (int) (((long)maxV * (long)maxH) % (long)(1e9+7));
    }
}