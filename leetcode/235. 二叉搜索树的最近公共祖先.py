# 给定一个二叉搜索树, 找到该树中两个指定节点的最近公共祖先。

# 百度百科中最近公共祖先的定义为：“对于有根树 T 的两个结点 p、q，最近公共祖先表示为一个结点 x，满足 x 是 p、q 的祖先且 x 的深度尽可能大（一个节点也可以是它自己的祖先）。”

# 例如，给定如下二叉搜索树:  root = [6,2,8,0,4,7,9,null,null,3,5]

#  

# 示例 1:

# 输入: root = [6,2,8,0,4,7,9,null,null,3,5], p = 2, q = 8
# 输出: 6
# 解释: 节点 2 和节点 8 的最近公共祖先是 6。
# 示例 2:

# 输入: root = [6,2,8,0,4,7,9,null,null,3,5], p = 2, q = 4
# 输出: 2
# 解释: 节点 2 和节点 4 的最近公共祖先是 2, 因为根据定义最近公共祖先节点可以为节点本身。

# 来源：力扣（LeetCode）
# 链接：https://leetcode-cn.com/problems/lowest-common-ancestor-of-a-binary-search-tree
# 著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
from LeetCode.tree import TreeNode


class Solution:
    def lowestCommonAncestor(self, root: 'TreeNode', p: 'TreeNode', q: 'TreeNode') -> 'TreeNode':
        self.ans = None
        if p == q:
            self.ans = p
        else:
            self.findLeaf(root, p, q)
        return self.ans

    def findLeaf(self, root, p, q):
        if root is None:
            return None
        else:
            left = self.findLeaf(root.left, p, q)
            right = self.findLeaf(root.right, p, q)

        if root == p:
            if left == q or right == q:
                self.ans = root
                return True
            else:
                return p
        elif root == q:
            if left == p or right == p:
                self.ans = root
                return True
            else:
                return q
        elif left is None:
            return right
        elif right is None:
            return left
        else:
            self.ans = root
            return True
