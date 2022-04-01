/**
 * @author Bao WJ
 * @date 2021/11/26
 * @link https://leetcode-cn.com/problems/search-in-a-binary-search-tree/
 */
package com.leetcode;

class Solution700 {
    public TreeNode searchBST(TreeNode root, int val) {
        if (root == null) {
            return null;
        }
        if (root.val < val) {
            return searchBST(root.right, val);
        } else if (root.val > val) {
            return searchBST(root.left, val);
        } else {
            return root;
        }
    }
}