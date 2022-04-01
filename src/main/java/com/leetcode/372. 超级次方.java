/**
 * @author Bao WJ
 * @date 2021/12/5
 * @link https://leetcode-cn.com/problems/super-pow/
 */
package com.leetcode;

class Solution372 {
    public int superPow(int a, int[] b) {
        int first = 0;
        int ans = 1;
        a %= 1337;
        while ((first = getFirst(first, b)) != -1) {
            if (divideTwo(first, b)) {
                ans = (ans * a) % 1337;
            }
            a = (a * a) % 1337;
        }
        return ans;
    }

    public boolean divideTwo(int first, int[] b) {
        int last = 0;
        for (int i = first; i < b.length; i++) {
            first = last * 10 + b[i];
            if (first % 2 == 0) {
                b[i] = first / 2;
                last = 0;
            } else {
                b[i] = (first - 1) / 2;
                last = 1;
            }
        }
        if (last == 1) {
            return true;
        }
        return false;
    }

    public int getFirst(int first, int[] B) {
        for (; first < B.length && B[first] == 0; first++) {}
        if (first < B.length && B[first] != 0) {
            return first;
        }
        return -1;
    }
}