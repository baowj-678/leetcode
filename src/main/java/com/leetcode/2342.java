package com.leetcode;

import java.util.HashMap;
import java.util.Map;

class Solution2342 {
    public int maximumSum(int[] nums) {
        HashMap<Integer, Integer[]> map = new HashMap<>();
        for (int num: nums) {
            int s = 0, tmp = num;
            while (tmp > 0) {
                s += tmp % 10;
                tmp /= 10;
            }

            if (map.containsKey(s)) {
                Integer[] integers = map.get(s);
                if (num > integers[1]) {
                    integers[0] = integers[1];
                    integers[1] = num;
                } else if (num > integers[0]) {
                    integers[0] = num;
                }
                map.put(s, integers);
            } else {
                map.put(s, new Integer[]{0, num});
            }
        }

        int res = -1;
        for (Map.Entry<Integer, Integer[]> entry: map.entrySet()) {
            if (entry.getValue()[0] != 0) {
                int s = entry.getValue()[0] + entry.getValue()[1];
                if (s > res) {
                    res = s;
                }
            }
        }
        return res;
    }

    public static void main(String[] args) {
        int[] nums = new int[]{18,43,36,13,7};
        Solution2342 so = new Solution2342();
        System.out.println(so.maximumSum(nums));
    }
}