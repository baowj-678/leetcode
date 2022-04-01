/**
 * @auhtor Bao WJ
 * @date 2021/12/6
 * @link https://leetcode-cn.com/problems/truncate-sentence/
 */
package com.leetcode;

class Solution1816 {
    public String truncateSentence(String s, int k) {
        for (int i = 0; i < s.length(); i++) {
            if (s.charAt(i) == ' ') {
                k--;
            }
            if (k == 0) {
                return s.substring(0, i);
            }
        }
        return s;
    }
}