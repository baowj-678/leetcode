/**
 * @author Bao WJ
 * @link https://leetcode-cn.com/problems/ransom-note/
 * @date 2021/12/4
 */
package com.leetcode;

class Solution383 {
    public boolean canConstruct(String ransomNote, String magazine) {
        int[] magNum = new int[26];
        int[] ranNum = new int[26];

        for (char c : ransomNote.toCharArray()) {
            ranNum[c - 'a'] ++;
        }
        for (char c : magazine.toCharArray()) {
            magNum[c - 'a'] ++;
        }

        for (int i = 0; i < 26; i++) {
            if (magNum[i] < ranNum[i]) {
                return false;
            }
        }
        return true;
    }
}