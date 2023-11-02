package com.leetcode;

import java.util.ArrayList;
import java.util.HashSet;

class Solution2103 {
    public int countPoints(String rings) {
        ArrayList<HashSet<Integer>> maps = new ArrayList<>();
        for (int i = 0; i < 10; i++) {
            maps.add(new HashSet<>());
        }
        for (int i = 0; 2 * i < rings.length(); i++) {
            int color = rings.charAt(2*i), idx = rings.charAt(2*i+1)-'0';
            maps.get(idx).add(color);
        }
        int res = 0;
        for (HashSet<Integer> map: maps) {
            if (map.size() >= 3) {
                res++;
            }
        }
        return res;
    }
}