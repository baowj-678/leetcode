/**
 * @author Bao WJ
 * @link https://leetcode-cn.com/problems/maximum-product-of-word-lengths/
 * @date 2021/11/17
 */
package com.leetcode;

class Solution318 {
    public int maxProduct(String[] words) {
        int[] flags = new int[words.length];
        for (int i = 0; i < words.length; i++) {
            String tmp = words[i];
            int flag = 0;
            for (char c : tmp.toCharArray()) {
                c -= 'a';
                flag |= (1 << c);
            }
            flags[i] = flag;
        }
        int res = 0;
        for (int i = 0; i < words.length; i++) {
            for (int j = i + 1; j < words.length; j++) {
                if ((flags[i] & flags[j]) == 0) {
                    res = Math.max(res, words[i].length() * words[j].length());
                }
            }
        }
        return res;
    }
}