/**
 * @author Bao WJ
 * @date 2021/11/18
 * @link https://leetcode-cn.com/problems/binary-tree-tilt/
 */
package com.leetcode;

class Solution563 {
    int sum = 0;

    public int findTilt(TreeNode root) {
        findSubTilt(root);
        return sum;
    }

    public int findSubTilt(TreeNode root) {
        if (root == null) {
            return 0;
        }
        int left = findSubTilt(root.left);
        int right = findSubTilt(root.right);
        sum += Math.abs(left - right);
        return left + right + root.val;
    }
}