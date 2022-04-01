/**
 * @author Bao WJ
 * @date 2021/11/27
 * @link https://leetcode-cn.com/problems/random-flip-matrix/
 */
package com.leetcode;

import java.util.*;

class Solution519 {
    Map<Integer, Integer> map = new HashMap<>();
    Random random = new Random();
    int z;
    int n;
    int total;

    public Solution519(int m, int n) {
        this.n = n;
        total = m*n;
        z = total;
    }

    public int[] flip() {
        int i = random.nextInt(z);
        int k = map.getOrDefault(i, i);
        int swap = map.getOrDefault(z-1, z-1); // arr[z-1]
        map.put(i, swap);
        map.put(z-1, k);
        --z;
        return new int[]{k/n, k%n};
    }

    public void reset() {
        z = total;
    }
}