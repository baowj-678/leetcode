package com.leetcode;

import java.util.HashMap;
import java.util.HashSet;

class Solution2586 {
    public int vowelStrings(String[] words, int left, int right) {
        int res = 0;
        HashSet<Character> vowel = new HashSet<>(5);
        vowel.add('a');
        vowel.add('e');
        vowel.add('i');
        vowel.add('o');
        vowel.add('u');
        for (int i = left; i <= right; i++) {
            String s = words[i];
            if (vowel.contains(s.charAt(0)) && vowel.contains(s.charAt(s.length()-1))) {
                res++;
            }
        }
        return res;
    }
}