# 给出集合 [1,2,3,…,n]，其所有元素共有 n! 种排列。

# 按大小顺序列出所有排列情况，并一一标记，当 n = 3 时, 所有排列如下：

# "123"
# "132"
# "213"
# "231"
# "312"
# "321"
# 给定 n 和 k，返回第 k 个排列。

# 说明：

# 给定 n 的范围是 [1, 9]。
# 给定 k 的范围是[1,  n!]。
# 示例 1:

# 输入: n = 3, k = 3
# 输出: "213"
# 示例 2:

# 输入: n = 4, k = 9
# 输出: "2314"

# 来源：力扣（LeetCode）
# 链接：https://leetcode-cn.com/problems/permutation-sequence
# 著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。


class Solution:
    def getPermutation(self, n: int, k: int) -> str:
        base = 1
        ans = ''
        n_ = n
        k -= 1
        nums = []
        for i in range(1, n + 1):
            base *= i
            nums.append(i)
        while n_ != 0:
            base //= n_
            n_ -= 1
            num = k // base
            ans = ans + str(nums[num])
            del(nums[num])
            k -= base * num
        return ans


if __name__ == "__main__":
    solution = Solution()
    print(solution.getPermutation(9, 89))
