package com.leetcode;

import javafx.util.Pair;

import java.util.ArrayList;
import java.util.Stack;

class Solution2736 {
    public int[] maximumSumQueries(int[] nums1, int[] nums2, int[][] queries) {
        ArrayList<Pair<Integer, Integer>> nums = new ArrayList<>(nums1.length);
        for (int i = 0; i < nums1.length; i++) {
            nums.add(new Pair<>(nums1[i], nums2[i]));
        }
        nums.sort((o1, o2) -> o2.getKey() - o1.getKey());

        ArrayList<Pair<int[], Integer>> queriesWithIdx = new ArrayList<>(queries.length);
        for (int i = 0; i < queries.length; i++) {
            queriesWithIdx.add(new Pair<>(queries[i], i));
        }
        queriesWithIdx.sort((o1, o2) -> o2.getKey()[0] - o1.getKey()[0]);
        Stack<Pair<Integer, Integer>> stack = new Stack<>();

        int[] res  = new int[queries.length];
        int idx = 0;
        for (Pair<int[], Integer> query: queriesWithIdx) {
            while (idx < nums.size() && nums.get(idx).getKey() >= query.getKey()[0]) {
                Pair<Integer, Integer> element = nums.get(idx);
                if (stack.empty()) {
                    stack.push(new Pair<>(element.getValue(), element.getValue() + element.getKey()));
                } else {
                    Pair<Integer, Integer> top = stack.peek();
                    if (element.getValue() > top.getKey()) {
                        while (element.getKey() + element.getValue() >= top.getValue()) {
                            stack.pop();
                            if (stack.empty()) {
                                break;
                            }
                            top = stack.peek();
                        }
                        stack.push(new Pair<>(element.getValue(), element.getValue() + element.getKey()));
                    }
                }
                idx++;
            }
            res[query.getValue()] = binarySearch(stack, query.getKey()[1]);
        }
        return res;
    }

    private int binarySearch(Stack<Pair<Integer, Integer>>stack, int v) {
        int bottom = 0, top = stack.size();
        while (bottom < top) {
            int mid = (bottom + top) / 2;
            if (stack.get(mid).getKey() < v) {
                bottom = mid + 1;
            } else {
                top = mid;
            }
        }
        if (top == stack.size()) {
            return -1;
        }
        return stack.get(top).getValue();
    }
}