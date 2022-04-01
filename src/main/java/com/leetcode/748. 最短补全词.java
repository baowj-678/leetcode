/**
 * @author BaoWJ
 * @link https://leetcode-cn.com/problems/shortest-completing-word/
 * @date 2021/12/10
 */
package com.leetcode;

import java.util.Arrays;

class Solution748 {
    public String shortestCompletingWord(String licensePlate, String[] words) {
        int[] plate = new int[26];
        int[] tmp = new int[26];
        int minLen = Integer.MAX_VALUE;
        String ans = null;
        Arrays.fill(plate, 0);
        for (char c : licensePlate.toLowerCase().toCharArray()) {
            if (Character.isAlphabetic(c)) {
                plate[c - 'a']++;
            }
        }
        Arrays.fill(tmp, 0);
        for (String word : words) {
            if (word.length() >= minLen) {
                continue;
            }
            for (char c : word.toCharArray()) {
                tmp[c - 'a']++;
            }
            boolean isOK = true;
            for (int i = 0; i < 26; i++) {
                if (tmp[i] < plate[i]) {
                    isOK = false;
                }
                tmp[i] = 0;
            }
            if (isOK) {
                minLen = word.length();
                ans = word;
            }
        }
        return ans;
    }
}