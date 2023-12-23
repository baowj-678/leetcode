package com.leetcode;

import javafx.beans.property.ReadOnlyIntegerWrapper;

import java.util.ArrayList;
import java.util.HashSet;
import java.util.Iterator;
import java.util.List;

class Solution1462 {
    public List<Boolean> checkIfPrerequisite(int numCourses, int[][] prerequisites, int[][] queries) {
        HashSet<Integer>[] preDirectSet = new HashSet[numCourses];
        HashSet<Integer>[] preSet = new HashSet[numCourses];
        for (int[] prerequisite: prerequisites) {
            if (preDirectSet[prerequisite[1]] == null) {
                preDirectSet[prerequisite[1]] = new HashSet<>();
            }
            preDirectSet[prerequisite[1]].add(prerequisite[0]);
        }

        for (int i = 0; i < numCourses; i++) {
            HashSet<Integer> tmp = new HashSet<>();
            dfs(i, preDirectSet, tmp, new boolean[numCourses]);
            preSet[i] = tmp;
        }

        List<Boolean> res = new ArrayList<>();
        for (int[] query: queries) {
            if (preSet[query[1]].contains(query[0])) {
                res.add(true);
            } else {
                res.add(false);
            }
        }
        return res;
    }

    private void dfs(int idx, HashSet<Integer>[] preDirectSet, HashSet<Integer> set, boolean[] visited) {
        if (visited[idx]) {
            return;
        }
        visited[idx] = true;
        set.add(idx);
        if (preDirectSet[idx] == null) {
            return;
        }
        for (int next: preDirectSet[idx]) {
            dfs(next, preDirectSet, set, visited);
        }
    }
}