package com.leetcode;

import java.util.PriorityQueue;

class Solution2530 {
    public long maxKelements(int[] nums, int k) {
        PriorityQueue<Integer> q = new PriorityQueue<>((o1, o2) -> {return -o1.compareTo(o2);});
        long s = 0;
        for (int num: nums) {
            q.add(num);
        }
        for (int i = 0; i < k; i++) {
            long peek = q.remove();
            s += peek;
            peek = (peek + 2) / 3;
            q.add((int) peek);
        }
        return s;
    }

    public static void main(String[] args) {
        Solution2530 solution2530 = new Solution2530();
        System.out.println(solution2530.maxKelements(new int[]{10,10,10,10,10}, 5));
    }
}