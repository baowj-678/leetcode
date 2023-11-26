package com.leetcode;

import java.util.Arrays;

class Solution2594 {
    public long repairCars(int[] ranks, int cars) {
        long l = 1, r = (long)Arrays.stream(ranks).min().getAsInt() * cars * cars;
        while (l < r) {
            long mid = (l + r) / 2;
            if (this.check(ranks, cars, mid)) {
                r = mid;
            } else {
                l = mid + 1;
            }
        }
        return l;
    }

    private boolean check(int[] ranks, int cars, long time) {
        for (int rank: ranks) {
            int car = (int) Math.floor(Math.sqrt((double)time / rank));
            cars -= car;
            if (cars <= 0) {
                return true;
            }
        }
        return cars <= 0;
    }
}