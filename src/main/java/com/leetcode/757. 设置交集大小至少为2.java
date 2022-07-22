package com.leetcode;

import java.util.Arrays;
import java.util.Comparator;
import java.util.Map;
import java.util.TreeMap;

class Solution757 {
    public int intersectionSizeTwo(int[][] intervals) {
        Arrays.sort(intervals, new Comparator<int[]>() {
            @Override
            public int compare(int[] o1, int[] o2) {
                if (o1[0] == o2[0]) {
                    if (o1[1] > o2[1]) {
                        return -1;
                    } else if (o1[1] == o2[1]) {
                        return 0;
                    } else {
                        return 1;
                    }
                } else if (o1[0] < o2[0]) {
                    return -1;
                } else {
                    return 1;
                }
            }
        });
        for (int[] interval : intervals) {
            System.out.println(interval[0] +"-"+interval[1]);
        }
        int first = Integer.MAX_VALUE, second = Integer.MAX_VALUE;
        int lastLeft = Integer.MAX_VALUE, ans = 0;
        for (int i = intervals.length-1; i >= 0; i--) {
            if (intervals[i][0] < lastLeft) {
                if (intervals[i][1] < first) {
                    ans += 2;
                    first = intervals[i][0];
                    second = intervals[i][0] + 1;
                } else if (intervals[i][1] < second) {
                    second = first;
                    first = intervals[i][0];
                    ans++;
                }
            }
        }
        return ans;
    }
}