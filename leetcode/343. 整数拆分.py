# 给定一个正整数 n，将其拆分为至少两个正整数的和，并使这些整数的乘积最大化。 返回你可以获得的最大乘积。

# 示例 1:

# 输入: 2
# 输出: 1
# 解释: 2 = 1 + 1, 1 × 1 = 1。
# 示例 2:

# 输入: 10
# 输出: 36
# 解释: 10 = 3 + 3 + 4, 3 × 3 × 4 = 36。
# 说明: 你可以假设 n 不小于 2 且不大于 58。

# 来源：力扣（LeetCode）
# 链接：https://leetcode-cn.com/problems/integer-break
# 著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。


class Solution:
    def integerBreak(self, n: int) -> int:
        if(n == 2):
            return 1
        elif(n == 3):
            return 2
        else:
            r = n // 3
            l = n - r * 3
            if(l == 1):
                r -= 1
                l += 3
            elif(l == 0):
                l = 1
            return 3 ** r * l


if __name__ == '__main__':
    solution = Solution()
    print(solution.integerBreak(67))