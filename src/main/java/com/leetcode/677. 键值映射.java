/**
 * @author Bao Wenjie
 * @link https://leetcode-cn.com/problems/map-sum-pairs/
 * @date 2921/11/14
 */
package com.leetcode;

import java.util.Arrays;

class MapSum {

    class Node {
        int val;
        String key;
        boolean isLeaf;
        Node[] next;

        public Node() {
            isLeaf = false;
            next = new Node[26];
            Arrays.fill(next, null);
        }
    }
    Node head;

    public MapSum() {
        head = new Node();
    }

    public void insert(String key, int val) {
        Node tmp = head;
        for (int i = 0; i < key.length(); i++) {
            int index = key.charAt(i) - 'a';
            if (tmp.next[index] == null) {
                Node node = new Node();
                tmp.next[index] = node;
            }
            tmp = tmp.next[index];
        }
        tmp.isLeaf = true;
        tmp.key = key;
        tmp.val = val;
    }

    public int sum(String prefix) {
        Node tmp = head;
        for (int i = 0; i < prefix.length(); i++) {
            int index = prefix.charAt(i) - 'a';
            if (tmp.next[index] == null) {
                return 0;
            }
            tmp = tmp.next[index];
        }
        return sumSubTree(tmp);
    }

    public int sumSubTree(Node root) {
        int sum = 0;
        if (root == null) {
            return 0;
        }
        if (root.isLeaf) {
            sum += root.val;
        }
        for (Node node : root.next) {
            if (node != null) {
                sum += sumSubTree(node);
            }
        }
        return sum;
    }
}

/**
 * Your MapSum object will be instantiated and called as such:
 * MapSum obj = new MapSum();
 * obj.insert(key,val);
 * int param_2 = obj.sum(prefix);
 */