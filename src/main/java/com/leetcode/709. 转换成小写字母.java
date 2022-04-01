/**
 * @author BaoWJ
 * @date 2021/12/12
 * @link https://leetcode-cn.com/problems/to-lower-case/
 */
package com.leetcode;


class Solution709 {
    public String toLowerCase(String s) {
        StringBuilder sb = new StringBuilder();
        s.chars().map(Character::toLowerCase).forEach((c) -> sb.append((char)c));
        return new String(sb);
    }
}