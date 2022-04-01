/**
 * @author Bao Wenjie
 * @link https://leetcode-cn.com/problems/shu-zu-zhong-shu-zi-chu-xian-de-ci-shu-lcof/
 * @date 2021/10/14
 */

package com.leetcode;

class Solution56 {
    public int[] singleNumbers(int[] nums) {
        int ans = 0;
        for(int x : nums) {
            ans ^= x;
        }
        int index = 0;
        while(ans != 0) {
            if((ans & 0b1) != 0) {
                break;
            }
            ans >>= 1;
            index++;
        }
        int num1 = 0, num2 = 0;
        for(int num : nums) {
            if (((num >> index) & 0b1) == 0) {
                num1 ^= num;
            } else {
                num2 ^= num;
            }
        }
        int[] a = {num1, num2};
        return a;
    }
}
