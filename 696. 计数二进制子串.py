# 给定一个字符串 s，计算具有相同数量0和1的非空(连续)子字符串的数量，并且这些子字符串中的所有0和所有1都是组合在一起的。

# 重复出现的子串要计算它们出现的次数。

# 示例 1 :

# 输入: "00110011"
# 输出: 6
# 解释: 有6个子串具有相同数量的连续1和0：“0011”，“01”，“1100”，“10”，“0011” 和 “01”。

# 请注意，一些重复出现的子串要计算它们出现的次数。

# 另外，“00110011”不是有效的子串，因为所有的0（和1）没有组合在一起。
# 示例 2 :

# 输入: "10101"
# 输出: 4
# 解释: 有4个子串：“10”，“01”，“10”，“01”，它们具有相同数量的连续1和0。
# 注意：

# s.length 在1到50,000之间。
# s 只包含“0”或“1”字符。

# 来源：力扣（LeetCode）
# 链接：https://leetcode-cn.com/problems/count-binary-substrings
# 著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。


class Solution:
    def countBinarySubstrings(self, s: str) -> int:
        continuousZeros, continuousOnes = 0, 0
        lastEle = ''
        cnt = 0
        for i in range(len(s)):
            if(s[i] == '0'):
                if(lastEle == '0'):
                    continuousZeros += 1
                else:
                    lastEle = '0'
                    continuousZeros = 1
                if(continuousZeros <= continuousOnes):
                    cnt += 1
            else:
                if(lastEle == '1'):
                    continuousOnes += 1
                else:
                    lastEle = '1'
                    continuousOnes = 1
                if(continuousOnes <= continuousZeros):
                    cnt += 1
        return cnt


if __name__ == '__main__':
    solution = Solution()
    s = "00100"
    print(solution.countBinarySubstrings(s))
