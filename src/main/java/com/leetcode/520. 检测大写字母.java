/**
 * @author Bao Wenjie
 * @link https://leetcode-cn.com/problems/detect-capital/
 * @date 2021/11/13
 */
package com.leetcode;

class Solution520 {
    public boolean detectCapitalUse(String word) {
        boolean isFirstLow = false;
        if (Character.isLowerCase(word.charAt(0))) {
            isFirstLow = true;
        }
        boolean isAfterLow = true;
        for (int i = 1; i < word.length(); i++) {
            if (i == 1) {
                if (Character.isLowerCase(word.charAt(i))) {
                    isAfterLow = true;
                } else {
                    isAfterLow = false;
                }
            } else if (Character.isLowerCase(word.charAt(i)) != isAfterLow) {
                return false;
            }
        }
        if (isFirstLow && !isAfterLow) {
            return false;
        }
        return true;
    }
}

