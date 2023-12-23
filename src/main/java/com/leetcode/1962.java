package com.leetcode;

import java.util.Comparator;
import java.util.PriorityQueue;

class Solution1962 {
    public int minStoneSum(int[] piles, int k) {
        int res = 0;
        PriorityQueue<Integer> pq = new PriorityQueue<>((Comparator.comparingInt(o -> -o)));
        for (int pile: piles) {
            pq.add(pile);
            res += pile;
        }
        for (int i = 0; i < k; i++) {
            int v = pq.poll();
            int minus = v/2;
            res -= minus;
            pq.add(v-minus);
        }
        return res;
    }
}