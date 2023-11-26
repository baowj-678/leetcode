package com.leetcode;

import com.sun.xml.internal.bind.v2.model.core.NonElement;
import sun.reflect.generics.tree.Tree;

import java.net.Inet4Address;
import java.util.HashMap;
import java.util.Map;

/**
 * Definition for a binary tree node.
 * public class TreeNode {
 *     int val;
 *     TreeNode left;
 *     TreeNode right;
 *     TreeNode() {}
 *     TreeNode(int val) { this.val = val; }
 *     TreeNode(int val, TreeNode left, TreeNode right) {
 *         this.val = val;
 *         this.left = left;
 *         this.right = right;
 *     }
 * }
 */
class Solution1457 {
    int count;
    public int pseudoPalindromicPaths (TreeNode root) {
        this.count = 0;
        dfs(root, new HashMap<>());
        return this.count;
    }

    private void dfs(TreeNode root, HashMap<Integer, Integer> map) {
        int v = 1;
        if (map.containsKey(root.val)) {
            v += map.get(root.val);
            map.replace(root.val, v);
        } else {
            map.put(root.val, v);
        }
        if (root.left != null) {
            dfs(root.left, map);
        }
        if (root.right != null) {
            dfs(root.right, map);
        }
        if (root.left == null && root.right == null) {
            int oddCnt = 0;
            for (Map.Entry<Integer, Integer> entry: map.entrySet()) {
                if (entry.getValue() % 2 == 1) {
                    oddCnt++;
                }
                if (oddCnt > 1) {
                     break;
                }
            }
            if (oddCnt <= 1) {
                this.count++;
            }
        }

        if (v == 1) {
            map.remove(root.val);
        } else {
            map.replace(root.val, v-1);
        }
    }
}