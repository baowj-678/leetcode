/**
 * @author Bao WJ
 * @link https://leetcode-cn.com/problems/integer-replacement/
 * @date 2021/11/19
 */
package com.leetcode;

import java.util.LinkedList;
import java.util.Queue;

class Solution397 {
    public int integerReplacement(int n) {
        Queue<Long> queue = new LinkedList<>();
        int cnt = 0;
        queue.add((long)n);
        queue.add(0L);
        while (!queue.isEmpty()) {
            long tmp =  queue.poll();
            if (tmp == 0) {
                queue.add(0L);
                cnt++;
            } else if (tmp == 1) {
                break;
            } else {
                if (tmp % 2 == 0) {
                    queue.add(tmp >> 1);
                } else {
                    queue.add(tmp + 1);
                    queue.add(tmp - 1);
                }
            }
        }
        return cnt;
    }
}