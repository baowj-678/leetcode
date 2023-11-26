package com.leetcode;


class Solution2760 {
    public int longestAlternatingSubarray(int[] nums, int threshold) {
        int res = 0;
        for (int l = 0; l < nums.length; l++) {
            boolean isEven = true;
            boolean isInclude = true;
            for (int i = l; i < nums.length; i++) {
                if (((nums[i] % 2 == 0) == isEven) && nums[i] <= threshold) {
                    isEven = !isEven;
                } else {
                    res = Math.max(res, i - l);
                    isInclude = false;
                    break;
                }
            }
            if (isInclude) {
                res = Math.max(res, nums.length - l);
            }

        }
        return res;
    }

    public static void main(String[] args) {
        int[] nums = new int[]{3,2,5,4};
        Solution2760 solution2760 = new Solution2760();
        System.out.println(solution2760.longestAlternatingSubarray(nums, 5));
    }
}