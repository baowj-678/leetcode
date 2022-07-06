# 给定一个二叉树，我们在树的节点上安装摄像头。

# 节点上的每个摄影头都可以监视其父对象、自身及其直接子对象。

# 计算监控树的所有节点所需的最小摄像头数量。

#  

# 示例 1：

# 输入：[0,0,null,0,0]
# 输出：1
# 解释：如图所示，一台摄像头足以监控所有节点。
# 示例 2：

# 输入：[0,0,null,0,null,0,null,null,0]
# 输出：2
# 解释：需要至少两个摄像头来监视树的所有节点。 上图显示了摄像头放置的有效位置之一。

# 提示：

# 给定树的节点数的范围是 [1, 1000]。
# 每个节点的值都是 0。

# 来源：力扣（LeetCode）
# 链接：https://leetcode-cn.com/problems/binary-tree-cameras
# 著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
from LeetCode.tree import TreeNode


class Solution:
    def minCameraCover(self, root: TreeNode) -> int:
        self.ans = 0
        a, b, c = self.dfs(root)
        return b

    def dfs(self, root):
        if root is None:
            return (float('inf'), 0, 0)
        else:
            la, lb, lc = self.dfs(root.left)
            ra, rb, rc = self.dfs(root.right)
            a = lc + rc + 1
            b = min(a, min(lb + ra, la + rb))
            c = min(a, lb + rb)
            return (a, b, c)
