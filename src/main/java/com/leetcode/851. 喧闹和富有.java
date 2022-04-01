/**
 * @author BaoWJ
 * @link https://leetcode-cn.com/problems/loud-and-rich/
 * @date 2021/12/15
 */
package com.leetcode;

import java.util.ArrayList;
import java.util.Arrays;
import java.util.List;

class Solution851 {
    int[] ans;
    int[] quiet;
    List<Integer>[] graph;

    public int[] loudAndRich(int[][] richer, int[] quiet) {
        this.quiet = quiet;
        this.ans = new int[quiet.length];
        Arrays.fill(this.ans, -1);
        this.graph = new ArrayList[quiet.length];
        for (int i = 0; i < graph.length; i++) {
            graph[i] = new ArrayList<>();
        }
        for (int[] tmp : richer) {
            graph[tmp[1]].add(tmp[0]);
        }
        for (int i = 0; i < this.ans.length; i++) {
            DFS(i);
        }
        return this.ans;
    }

    public void DFS(int node) {
        if (this.ans[node] != -1) {
            return;
        }
        int minV = this.quiet[node];
        int minP = node;
        int tmpV, tmpP;
        for (int x : this.graph[node]) {
            tmpP = this.ans[x];
            if (tmpP == -1) {
                DFS(x);
                tmpP = this.ans[x];
            }
            tmpV = this.quiet[tmpP];
            if (tmpV < minV) {
                minV = tmpV;
                minP = tmpP;
            }
        }
        this.ans[node] = minP;
    }
}